package sql

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Import du pilote SQLite
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))

	data := Employees{}
	//employees := getEmails()

	if r.Method == http.MethodPost {
		idFormulaireEmployee := r.FormValue("employee")
		idFormulaireManager := r.FormValue("employeeManager")
		idFormulaireHIre := r.FormValue("hire")
		idFormulaireIncrease := r.FormValue("increase")
		print(idFormulaireEmployee, idFormulaireHIre, idFormulaireManager, idFormulaireIncrease)
		postId := idFormulaireEmployee[0]
		isPresent := idFormulaireEmployee[1]
		employees := recoveryEmployees(postId, isPresent)
		managers := recoveryManagers(managerId, employeeId)
		posts := recoveryPosts(postId)
		departemnts := recoveryDepartements(id)
		fmt.Println(employees, managers, posts, departemnts)
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}

	//tmpl.Execute(w, struct{ Employees []string }{Employees: employees})
}
