package sql

import (
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Import du pilote SQLite
)



func Home(w http.ResponseWriter, r *http.Request) {
tmpl := template.Must(template.ParseFiles("templates/Home.html"))

	data := Employees{}
	allEmploye := recuperationEmployee()
	if r.Method == http.MethodPost{
		idFormulaireEmployee := r.FormValue("employee")
		idFormulaireManager := r.FormValue("employeeManager") 
		idFormulaireHIre := r.FormValue("hire")
		idFormulaireIncrease := r.FormValue("increase")
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
