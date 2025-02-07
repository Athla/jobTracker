package main

import (
	"jobTracker/internal/database"
	"os"

	"github.com/charmbracelet/log"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	db := database.New()

	user := os.Getenv("ADMIN_USERNAME")
	if user == "" {
		log.Fatalf("Error: ADMIN_USERNAME is a required enviroment variable.")
		os.Exit(1)
	}
	pwd := os.Getenv("ADMIN_PASSWORD")
	if pwd == "" {
		log.Fatalf("Error: ADMIN_PWD is a required enviroment variable.")
		os.Exit(1)
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Unable to generate pwd hash due: %s", err)
		os.Exit(1)
	}

	_, err = db.Exec(`
        insert into admin_user (username, password_hash)
        values (?, ?)
        on conflict(username) do update set password_hash = ?`,
		user, hashedPwd, hashedPwd)
	if err != nil {
		log.Fatalf("Unable to setup admin user correctly due: %s", err)
		os.Exit(1)
	}

	log.Info("Admin user correctly setup")
}
