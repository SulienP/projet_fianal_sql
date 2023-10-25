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

func getSubordinates(idManager string) []Employees {
	db := getDataBase()

	defer db.Close()

	rows, errQuery := db.Query("SELECT firstName, lastName FROM employees INNER JOIN posts ON employees.postId = posts.postId WHERE posts.managerId = " + idManager)
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
