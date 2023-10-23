package server

import (
	"fmt"
	"html/template"
	"net/http"
)

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

		login := r.FormValue("login")
		password := r.FormValue("password")
		email := r.FormValue("email")
		fullname := r.FormValue("fullname")
		age := r.FormValue("age")
		car := r.FormValue("car")

		query := `INSERT INTO users (login, password, email, fullname, age, car) VALUES ($1, $2, $3, $4, $5, $6)`

		_, err = app.DB.Exec(query, login, password, email, fullname, age, car)
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

		query := `DELETE FROM users WHERE login = $1 `
		_, err = app.DB.Exec(query, login)
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
