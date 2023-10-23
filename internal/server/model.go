package server

import (
	"database/sql"
)

type Err struct {
	Text string
	Code int
}

type App struct {
	DB *sql.DB
}
