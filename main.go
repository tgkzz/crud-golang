package main

import (
	"crudapp/internal/database"
	"crudapp/internal/server"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db, err := database.DBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := &server.App{DB: db}

	log.Print("Server starting on http://localhost:4000/")
	err = http.ListenAndServe(":4000", server.NewServer(app))
	if err != nil {
		log.Fatal(err)
	}

}
