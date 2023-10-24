package database

import (
	"crudapp/internal/models"
	"database/sql"
)

func ReadFromDb(db *sql.DB, login string) (models.User, error) {
	Result := models.User{}

	query := `SELECT * FROM users WHERE login = $1`
	err := db.QueryRow(query, login).Scan(&Result.Login, &Result.Password, &Result.Fullname, &Result.Email, &Result.Age, &Result.Car)
	if err != nil {
		return Result, err
	}
	return Result, nil
}
