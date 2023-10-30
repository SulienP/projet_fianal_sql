package sql

import (
	"log"
	"net/http"
	"text/template"
)

func Login(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))

	data:= Employees{}
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}