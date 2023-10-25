package sql

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/Home.html"))

	// Get the list of employees
	employees := getAllEmployees()

	if r.Method == http.MethodPost {
		idFormulaireEmployee := r.FormValue("employee")
		idFormulaireManager := r.FormValue("employeeManager")
		idFormulaireHire := r.FormValue("hire")
		idFormulaireIncrease := r.FormValue("increase")
		fmt.Println(idFormulaireEmployee, idFormulaireHire, idFormulaireManager, idFormulaireIncrease)

		// Handle form submissions here if needed
	}

	// Pass the list of employees to the template
	err := tmpl.Execute(w, employees)
	if err != nil {
		log.Fatal(err)
	}
}
