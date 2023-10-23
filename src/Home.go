package sql

import (
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Import du pilote SQLite
)



func Home(w http.ResponseWriter, r *http.Request) {
tmpl := template.Must(template.ParseFiles("templates/Home.html"))

	data := Employees{
		
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
