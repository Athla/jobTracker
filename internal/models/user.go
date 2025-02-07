package models

import (
	"database/sql"
	"errors"

	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) HashPwd() error {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		// This shouldn't ever happen
		log.Errorf("Unable to hash pwd due: %v", err)
		return err
	}

	u.Password = string(hashedPwd)
	return nil
}

func (u *User) CheckPwd(entryPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(entryPwd))
}

func CreateUser(db *sqlx.DB, username, pwd string) error {
	stmt, err := db.Preparex("INSERT INTO users (username, password) VALUES (?, ?)")
	if err != nil {
		log.Errorf("Unable to prepare statement due: %v", err)
		return err
	}

	defer stmt.Close()

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		// This shouldn't ever happen
		log.Errorf("Unable to hash pwd due: %v", err)
		return err
	}

	_, err = stmt.Exec(username, hashedPwd)
	if err != nil {
		log.Errorf("Unable to create user due: %v", err)
		return err
	}

	return nil
}

func GetUserByUsername(db *sqlx.DB, username string) (*User, error) {
	row := db.QueryRowx("SELECT id, username, password FROM users where username = ?", username)

	var user User
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("User not found!")
		}
		log.Errorf("Unable to get user due: %v", err)
		return nil, err
	}

	return &user, nil
}
