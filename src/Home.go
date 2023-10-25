package sql

import (
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
		idEmployee := r.FormValue("employee")
		idManager := r.FormValue("manager")
		idHire := r.FormValue("hire")
		idIncrease := r.FormValue("increase")
		print(idEmployee, idManager, idHire, idIncrease)

		/*
			postId := idEmployee[0]
			isPresent := idEmployee[1]
			employees := recoveryEmployees(postId, isPresent)
			managers := recoveryManagers(managerId, employeeId)
			posts := recoveryPosts(postId)
			departemnts := recoveryDepartements(id)
			fmt.Println(employees, managers, posts, departemnts)
		*/
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}

	//tmpl.Execute(w, struct{ Employees []string }{Employees: employees})
}
