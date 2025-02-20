package database

import (
	"embed"
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
)

//go:embed migrations/001_core.sql
var migrationFiles embed.FS

func RunMigrations(db *sqlx.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`)

	if err != nil {
		log.Errorf("Failed to create migrations table: %s", err)
		return err
	}

	applied := make(map[string]bool)
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

		applied[name] = true
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
		if applied[migration] {
			log.Infof("Migration %s already applied.", migration)
			continue
		}

		log.Infof("Applying migrations '%s'", migration)
		if err := applyMigration(db, migration); err != nil {
			log.Errorf("Unable to apply migration due: %s", err)
			return err
		}
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

	log.Infof("Rolling back migration '%s'", lastMigration)
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
	defer tx.Rollback()

	_, err = tx.Exec(downMigration)
	if err != nil {
		log.Errorf("Unable to execute transaction due: %s", err)
		return err
	}

	_, err = tx.Exec("DELETE FROM migrations WHERE name = ?", lastMigration)
	if err != nil {
		log.Errorf("Unable to remove migration '%s' from database due: %s", lastMigration, err)
		return err
	}

	log.Infof("Rolled back migration %s", lastMigration)

	return tx.Commit()
}

func RollbackAllMigrations(db *sqlx.DB) error {
	for {
		var count int
		err := db.Get(&count, "SELECT COUNT(*) FROM migrations")
		if err != nil {
			return err
		}
		if count == 0 {
			break
		}

		if err := RollbackMigration(db); err != nil {
			return err
		}
	}
	return nil
}

func applyMigration(db *sqlx.DB, migration string) error {
	content, err := migrationFiles.ReadFile(filepath.Join("migrations", migration))
	if err != nil {
		log.Errorf("Unable to read '%s' due: %s", filepath.Join("migrations", migration), err)
		return err
	}

	parts := strings.Split(string(content), "-- Down")
	if len(parts) != 2 {
		return fmt.Errorf("invalid migration format in %s", migration)
	}

	upSQL := strings.Split(parts[0], "-- Up")[1]

	tx, err := db.Beginx()
	if err != nil {
		log.Errorf("Unable to start transaction due: %s", err)
		return err
	}
	defer tx.Rollback()

	if _, err := tx.Exec(strings.TrimSpace(upSQL)); err != nil {
		log.Errorf("Unable to execute migration due: %s", err)
		return err
	}

	if _, err := tx.Exec("INSERT INTO migrations (name) VALUES (?)", migration); err != nil {
		log.Errorf("Unable to insert migration '%s' to database due: %s", migration, err)
		return err
	}

	return tx.Commit()
}
