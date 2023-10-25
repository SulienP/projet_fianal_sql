package sql

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))

	if r.Method == http.MethodPost {
		employeeId := r.FormValue("employeeId")

		fired(employeeId)
	}

	currentData := getAllEmployees()

	err := tmpl.Execute(w, currentData)
	if err != nil {
		log.Fatal(err)
	}
}
