package sql

import (
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))

	// Get the list of employees
	//currentData := getAllEmployees()
	currentData := getEmployees()

	/*
		if r.Method == http.MethodPost {
			idEmployee := r.FormValue("employee")
			idManager := r.FormValue("employeeManager")
			idHire := r.FormValue("hire")
			idIncrease := r.FormValue("increase")
			fmt.Println(idEmployee, idManager, idHire, idIncrease)

			// Handle form submissions here if needed
		}
	*/

	// Pass the list of employees to the template
	err := tmpl.Execute(w, currentData)
	if err != nil {
		log.Fatal(err)
	}
}
