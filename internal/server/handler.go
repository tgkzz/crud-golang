package server

import (
	"crudapp/internal/database"
	"crudapp/internal/models"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

type Err struct {
	Text string
	Code int
}

type App struct {
	DB *sql.DB
}

func ErrorHandler(w http.ResponseWriter, code int) {
	w.WriteHeader(code)

	tmpl, err := template.ParseFiles("templates/html/error.html")
	if err != nil {
		text := fmt.Sprintf("Error 500\n Oppss! %s", http.StatusText(http.StatusInternalServerError))
		http.Error(w, text, http.StatusInternalServerError)
		return
	}

	res := &Err{Text: http.StatusText(code), Code: code}
	err = tmpl.Execute(w, &res)
	if err != nil {
		text := fmt.Sprintf("Error 500\n Oppss! %s", http.StatusText(http.StatusInternalServerError))
		http.Error(w, text, http.StatusInternalServerError)
		return
	}
}

func (app *App) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("templates/html/index.html")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	default:
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

}

func (app *App) CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/create" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("templates/html/create.html")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	case "POST":
		tmpl, err := template.ParseFiles("templates/html/success.html")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		user := models.User{
			Login:    r.FormValue("login"),
			Password: r.FormValue("password"),
			Email:    r.FormValue("email"),
			Fullname: r.FormValue("fullname"),
			Age:      r.FormValue("age"),
			Car:      r.FormValue("car"),
		}

		if err := database.InsertToDB(app.DB, user); err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	default:
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
}

func (app *App) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/delete" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("templates/html/delete.html")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	case "POST":
		tmpl, err := template.ParseFiles("templates/html/success.html")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		login := r.FormValue("login")
		if login == "" {
			ErrorHandler(w, http.StatusBadGateway)
			return
		}

		if err = database.DeleteInDb(login, app.DB); err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	default:
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
}

func (app *App) ReadHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/read" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("templates/html/read.html")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	case "POST":
		login := r.FormValue("login")

		user, err := database.ReadFromDb(app.DB, login)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("templates/html/readAnswer.html")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		if err = tmpl.Execute(w, user); err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	}
}
