package sql

import (
	"fmt"
	"database/sql"
	"log"
)

var db *sql.DB

func init() {
	// Open the database connection during initialization
	var errSQLOpen error
	db, errSQLOpen = sql.Open("sqlite3", "public/company.db")
	if errSQLOpen != nil {
		log.Fatal(errSQLOpen)
	}
}

func getEmails() []string {
	rows, errQuery := db.Query("SELECT email FROM employees")
	if errQuery != nil {
		log.Fatal(errQuery)
	}
	defer rows.Close()

	var emails []string

	for rows.Next() {
		var email string
		err := rows.Scan(&email)
		if err != nil {
			log.Fatal(err)
		}
		emails = append(emails, email)
	}

	return emails
}

// type Employee struct {
// 	EmployeeID   int
// 	PostID       int
// 	FirstName    string
// 	LastName     string
// 	Email        string
// 	Password     string
// 	IsPresent    bool
// 	Salary       int
// 	Schedule     string
// 	BreakTimes   string
// 	DateHire     string
// 	EndContract  string
// }

func getAllEmployees() []Employees {
	database, _ := sql.Open("sqlite3", "./public/compagny.db")
	defer database.Close()
	rows, errQuery := database.Query("SELECT * FROM employees")
	if errQuery != nil {
		log.Fatal(errQuery)
	}
	defer rows.Close()

	var employees []Employees

	for rows.Next() {
		var employee Employees
		err := rows.Scan(&employee.EmployeeId, &employee.PostId, &employee.FirstName, &employee.LastName, &employee.Email, &employee.Password, &employee.IsPresent, &employee.Salary, &employee.Schedule, &employee.BreackTimes, &employee.DateHire, &employee.EndContract)
		if err != nil {
			log.Fatal(err)
		}

		employees = append(employees, employee)
	}
	fmt.Println(employees)
	return employees
}
