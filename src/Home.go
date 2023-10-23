package sql

import (
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Import du pilote SQLite
)

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/Home.html"))

	employees := getEmails()

	t.Execute(w, struct{ Employees []string }{Employees: employees})
}
