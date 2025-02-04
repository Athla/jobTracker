package database

import (
	"embed"
	"path/filepath"
	"sort"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func RunMigrations(db *sqlx.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		name TEXT NOT NULL,
		applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`)

	if err != nil {
		log.Errorf("Failed to create migrations table: %s", err)
		return err
	}

	appliedMigrations := make(map[string]bool)
	rows, err := db.Query("SELECT name FROM migrations")
	if err != nil {
		log.Errorf("Unable to get transactions due: %s", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Errorf("Unable to scan row due: %s", err)
			return err
		}

		appliedMigrations[name] = true
	}

	entries, err := migrationFiles.ReadDir("migrations")
	if err != nil {
		log.Errorf("Unable to read directory due: %s", err)
		return err
	}

	var migrations []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".sql") {
			migrations = append(migrations, entry.Name())
		}
	}

	sort.Strings(migrations)

	for _, migration := range migrations {
		if appliedMigrations[migration] {
			log.Infof("Migration %s already applied.", migration)
			continue
		}

		tx, err := db.Beginx()
		if err != nil {
			log.Errorf("Unable to start transaction due: %s", err)
			return err
		}

		content, err := migrationFiles.ReadFile(filepath.Join("migrations", migration))
		if err != nil {
			tx.Rollback()
			log.Errorf("Unable to read migration due: %s", err)
			return err
		}

		sections := strings.Split(string(content), "-- Down")
		if len(sections) != 2 {
			tx.Rollback()
			log.Errorf("Migration file %s has a invalid format.", migration)
			return err
		}

		upMigration := strings.Split(sections[0], "-- Up")[1]

		_, err = tx.Exec(upMigration)
		if err != nil {
			tx.Rollback()
			log.Errorf("Unable to execute migration %s due: %s", migration, err)
			return err
		}

		_, err = tx.Exec("INSERT INTO migrations (name) VALUES(?)", migration)
		if err != nil {
			tx.Rollback()
			log.Errorf("Unable to add migration '%s 'to records due: %s", migration, err)
			return err
		}

		if err := tx.Commit(); err != nil {
			log.Error("Unable to commit transaction due: %s", err)
			return err
		}

		log.Infof("Migration '%s' applied", migration)
	}

	return nil
}

func RollbackMigration(db *sqlx.DB) error {
	var lastMigration string
	err := db.QueryRowx(`SELECT name FROM migrations ORDER BY applied_at DESC LIMIT 1`).Scan(&lastMigration)

	if err != nil {
		log.Errorf("Unable to get last migration due: %s", err)
		return err
	}

	cnt, err := migrationFiles.ReadFile(filepath.Join("migrations", lastMigration))
	if err != nil {
		log.Errorf("Failed to read migration file '%s' due: %s", lastMigration, err)
		return err
	}

	sections := strings.Split(string(cnt), "-- Down")
	if len(sections) != 2 {
		log.Error("Invalid formation format")
		return err
	}

	downMigration := sections[1]

	tx, err := db.Beginx()
	if err != nil {
		log.Errorf("Unable to begin transaction due: %s", err)
		return err
	}

	_, err = tx.Exec(downMigration)
	if err != nil {
		tx.Rollback()
		log.Errorf("Unable to execute transaction due: %s", err)
		return err
	}

	_, err = tx.Exec("DELTE FROM migrations WHERE name = ?", lastMigration)
	if err != nil {
		tx.Rollback()
		log.Errorf("Unable to remove migration '%s' from database due: %s", lastMigration, err)
		return err
	}

	log.Infof("Rolled back migration %s", lastMigration)

	return nil
}
