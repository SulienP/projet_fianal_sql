package sql

import (
	"database/sql"
	"fmt"
	"log"
)

func getDataBase() *sql.DB {
	db, errSQLOpen := sql.Open("sqlite3", "public/compagny.db")
	if errSQLOpen != nil {
		fmt.Println(errSQLOpen)
	}

	return db
}

func getEmployees() []Employees {
	db := getDataBase()

	defer db.Close()

	rows, errQuery := db.Query("SELECT * FROM employees")
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
	return employees
}

func getColleague(nameService string) []Employees {
	db := getDataBase()

	defer db.Close()

	rows, errQuery := db.Query("SELECT employees.firstName, employees.lastName FROM employees INNER JOIN departements ON employees.postId = departements.departementId WHERE departements.serviceName = '" + nameService + "'")
	if errQuery != nil {
		log.Fatal(errQuery)
	}
	defer rows.Close()

	var employees []Employees

	for rows.Next() {
		var employee Employees
		err := rows.Scan(&employee.FirstName, &employee.LastName)
		if err != nil {
			log.Fatal(err)
		}

		employees = append(employees, employee)
	}
	return employees
}
