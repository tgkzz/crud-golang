package database

import (
	"crudapp/internal/models"
	"database/sql"
	"errors"
)

func InsertToDB(db *sql.DB, user models.User) error {
	query := `INSERT INTO users (login, password, email, fullname, age, car) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.Exec(query, user.Login, user.Password, user.Email, user.Fullname, user.Age, user.Car)
	if err != nil {
		return errors.New("error in inserting data to user")
	}
	return nil
}
