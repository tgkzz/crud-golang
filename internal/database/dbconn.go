package database

import "database/sql"

var connStr string = "postgresql://postgres:1234@localhost:5432/crudgolang?sslmode=disable"

func DBConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
