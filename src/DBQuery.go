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

func getAllEmployees() []Employees {
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


func addEmployees(postId int, firstName string, lastName string, email string, password string, isPresent string, salary int, schedule string, breakTimes string, dateHire string, endContract string) {
    fmt.Println("ici", postId, firstName, lastName, email, password, isPresent, salary, schedule, breakTimes, dateHire, endContract)
    
    db := getDataBase()
    defer db.Close()

    insertStatement := "INSERT INTO employees (postId, firstName, lastName, email, password, isPresent, salary, schedule, breakTimes, dateHire, endContract) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
    stmt, err := db.Prepare(insertStatement)
    fmt.Println(insertStatement)
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(postId, firstName, lastName, email, password, isPresent, salary, schedule, breakTimes, dateHire, endContract)
    if err != nil {
        log.Fatal(err)
    }
}
        var postName string

func getPost(employeeID int) string{
    db := getDataBase()
    defer db.Close()

    selectStatement := "SELECT posts.postName FROM posts  INNER JOIN employees ON posts.postId = employees.postId WHERE employees.employeeId = ?"
    stmt, err := db.Prepare(selectStatement)
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    rows, err := stmt.Query(employeeID)
    if err != nil {
        log.Fatal(err)
    }

    defer rows.Close()

    for rows.Next() {
        if err := rows.Scan(&postName); err != nil {
            log.Fatal(err)
        }
    }

    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }

	return postName
}

func newSalary(employeeID int, newSalary int) {
    db := getDataBase()
    defer db.Close()

    updateSalary := "UPDATE employees SET salary = ? WHERE employeeId = ?"
    stmt, err := db.Prepare(updateSalary)
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(newSalary, employeeID)
    if err != nil {
        log.Fatal(err)
    }
}
