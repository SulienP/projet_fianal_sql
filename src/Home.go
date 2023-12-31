package sql

import (

	"html/template"
	"log"
	"net/http"
	"strconv"
)

type EmployeeData struct {
    EmployeesList   []Employees
    OneEmployee     Employees
    PostName        string
	Manager 		string
	Departement	string
} 

var employeData EmployeeData

func Home(w http.ResponseWriter, r *http.Request) {	
	    tmpl := template.Must(template.ParseFiles("templates/home.html"))

	if r.Method == http.MethodPost {
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
        employeeData := Employees{
             
            PostId:       value,
            FirstName:    FirstName,
            LastName:     LastName,
            Email:        Email,
            Password:     Password,
            IsPresent:    IsPresent,
            Salary:       salary,
            Schedule:     Schedule,
            BreackTimes:  BreakTimes,
            DateHire:     DateHire,
            EndContract:  EndContract,
        }

        if employeeData.FirstName != "" {
            addEmployees(employeeData.PostId, employeeData.FirstName, employeeData.LastName, employeeData.Email, employeeData.Password, employeeData.IsPresent, employeeData.Salary, employeeData.Schedule, employeeData.BreackTimes, employeeData.DateHire, employeeData.EndContract)
        }
    }
}
    employeesList := getAllEmployees()
    employeData.EmployeesList = employeesList

    if r.Method == http.MethodPost {
        idEmploye := r.FormValue("idEmploye")
        id := 0
        if idEmploye != "" {
            id, _ = strconv.Atoi(idEmploye)
            oneEmployee := getEmployeData(id)
			getManager := getManagers((id))
			getDepartement := getDepartementNames(id)
            namePost := getPost(id)
            employeData.OneEmployee = oneEmployee
            employeData.PostName = namePost
			employeData.Manager = getManager
			employeData.Departement = getDepartement
        }
    }
    if r.Method == http.MethodPost {
        update_salary:= r.FormValue("new_salary")
        if update_salary != ""{
            
        update_salarytoi, _:=strconv.Atoi(update_salary)
        id := r.FormValue("id")
        idtoi,_:=strconv.Atoi(id)
        newSalary(idtoi,update_salarytoi)

        }
    }
    toFired := r.FormValue("fired")
	if toFired != "" {
		fired(toFired)
	}
    err := tmpl.Execute(w, employeData)
    if err != nil {
        log.Fatal(err)
    }


}
