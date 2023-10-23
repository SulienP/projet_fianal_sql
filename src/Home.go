package sql

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // Import du pilote SQLite
)



func Home(w http.ResponseWriter, r *http.Request) {
tmpl := template.Must(template.ParseFiles("templates/Home.html"))

	data := Employees{}

	
	if r.Method == http.MethodPost{
		idFormulaireEmployee := r.FormValue("employee")
		
		idFormulaireManager := r.FormValue("employeeManager") 
		idFormulaireHIre := r.FormValue("hire")
		idFormulaireIncrease := r.FormValue("increase")
	}
		allEmploye := RecuperationEmployee(postid  ,  isPresent )
	allManager := RecuperationManager( managerId, employeeId)
	allPost := RecuperationPost( postId )
	allDepartemnts := RecuperationDepartement(id )
	fmt.Println(allEmploye, allManager, allPost, allDepartemnts)
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
