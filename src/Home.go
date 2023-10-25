package sql

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)


type EmployeeData struct {
	PostID      int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	IsPresent   string
	Salary      int
	Schedule    string
	BreakTimes  string
	DateHire    string
	EndContract string
	PostName    string
}

var employeData []EmployeeData

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
<<<<<<< HEAD
    var namePost string
	if r.Method == http.MethodPost {
		PostID := r.FormValue("postId")
		FirstName := r.FormValue("firstName")
		LastName := r.FormValue("lastName")
		Email := r.FormValue("email")
		Password := r.FormValue("password")
		IsPresent := r.FormValue("isPresent")
		Salary := r.FormValue("salary")
		Schedule := r.FormValue("schedule")
		BreakTimes := r.FormValue("breakTimes")
		DateHire := r.FormValue("dateHire")
		EndContract := r.FormValue("endContract")
		value, _ := strconv.Atoi(PostID)
		salary, _ := strconv.Atoi(Salary)
		employeeData := EmployeeData{
			PostID:      value,
			FirstName:   FirstName,
			LastName:    LastName,
			Email:       Email,
			Password:    Password,
			IsPresent:   IsPresent,
			Salary:      salary,
			Schedule:    Schedule,
			BreakTimes:  BreakTimes,
			DateHire:    DateHire,
			EndContract: EndContract,
			PostName:    namePost,
		}

		employeData = append(employeData, employeeData)
		if employeData[0].FirstName != "" {
			addEmployees(employeData[0].PostID, employeData[0].FirstName, employeData[0].LastName, employeData[0].Email, employeData[0].Password, employeData[0].IsPresent, employeData[0].Salary, employeData[0].Schedule, employeData[0].BreakTimes, employeData[0].DateHire, employeData[0].EndContract)
		}
	}

	employees := getAllEmployees()
	if r.Method == http.MethodPost {
		idEmploye := r.FormValue("idEmploye")
		if idEmploye != "" {
			id, _ := strconv.Atoi(idEmploye)

			namePost := getPost(id)
            employees[0].PostName = namePost
		}
	}
	err := tmpl.Execute(w, employees)
	if err != nil {
		log.Fatal(err)
	}
    
=======

	// Get the list of employees
	//currentData := getEmployees()
	currentData := getColleague("Employees")

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
>>>>>>> 579b9947785c0b78517dc92661b688dd71ba0009
}
