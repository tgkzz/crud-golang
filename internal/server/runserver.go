package server

import (
	"net/http"
)

func NewServer(app *App) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/css"))))

	mux.HandleFunc("/", app.IndexHandler)
	mux.HandleFunc("/create", app.CreateHandler)
	mux.HandleFunc("/delete", app.DeleteHandler)

	return mux
}
