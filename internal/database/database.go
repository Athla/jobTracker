package database

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type Service interface {
	Health() map[string]string

	Close() error
}

type service struct {
	db *sqlx.DB
}

var (
	dburl = os.Getenv("BLUEPRINT_DB_URL")
)

func New() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", dburl)
	if err != nil {
		log.Fatalf("Unable to connect to database due: %v", err)
	}

	return db
}
