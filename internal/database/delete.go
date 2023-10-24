package database

import (
	"database/sql"
	"errors"
)

func DeleteInDb(login string, db *sql.DB) error {
	if login == "" {
		return errors.New("login must be not empty")
	}

	query := `DELETE FROM users WHERE login = $1`
	_, err := db.Exec(query, login)
	if err != nil {
		return errors.New("cannot delete user with specified login")
	}

	return nil
}
