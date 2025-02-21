package main

import (
	"jobTracker/internal/database"
	"os"

	"github.com/charmbracelet/log"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	db := database.New()
	defer db.Close()

	user := os.Getenv("ADMIN_USERNAME")
	if user == "" {
		user = "admin"
		log.Error("Error: ADMIN_USERNAME is a required enviroment variable.")
	}
	pwd := os.Getenv("ADMIN_PASSWORD")
	if pwd == "" {
		pwd = "adminpwd"
		log.Error("Error: ADMIN_PWD is a required enviroment variable.")
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Unable to generate pwd hash due: %s", err)
		os.Exit(1)
	}

	_, err = db.Exec(`
        INSERT INTO admin_user (username, password_hash)
        VALUES (?, ?)
        ON CONFLICT(username) DO UPDATE SET PASSWORD_HASH = ?`,
		user, hashedPwd, hashedPwd)
	if err != nil {
		log.Fatalf("Unable to setup admin user correctly due: %s", err)
		os.Exit(1)
	}

	log.Info("Admin user correctly setup")
}
