package main

import (
	"flag"
	"jobTracker/internal/database"
	"os"

	"github.com/charmbracelet/log"
)

func main() {
	down := flag.Bool("down", false, "Rollback migrations")
	all := flag.Bool("all", false, "Rollback all migrations, only with -down")

	db := database.New()

	defer db.Close()

	if *down {
		if *all {
			log.Info("Running back all migrations...")
			if err := database.RollbackAllMigrations(db); err != nil {
				log.Errorf("Unable to rollbback migrations due: %s", err)
			}
			log.Info("Rolledback all migrations")
			os.Exit(1)
		} else {
			log.Info("Rolling back last migration...")
			if err := database.RollbackAllMigrations(db); err != nil {
				log.Errorf("Unable to rollback migrations due: %s", err)
				os.Exit(1)
			}
			log.Info("Rolled back last migration with success.")
		}
	} else {
		log.Info("Running migrations")
		if err := database.RunMigrations(db); err != nil {
			log.Errorf("Unable to run migrations due: %s", err)
			os.Exit(1)
		}
		log.Info("Successfully ran migrations.")
	}
}
