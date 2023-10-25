package sql

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))

	// Get the list of employees
	employees := getAllEmployees()

	if r.Method == http.MethodPost {
		idEmployee := r.FormValue("employee")
		idManager := r.FormValue("employeeManager")
		idHire := r.FormValue("hire")
		idIncrease := r.FormValue("increase")
		fmt.Println(idEmployee, idManager, idHire, idIncrease)

		// Handle form submissions here if needed
	}

	// Pass the list of employees to the template
	err := tmpl.Execute(w, employees)
	if err != nil {
		log.Fatal(err)
	}
}
