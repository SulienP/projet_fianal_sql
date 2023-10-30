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

func getEmails() []string {
	db := getDataBase()

	rows, errQuery := db.Query("SELECT email FROM employees")
	if errQuery != nil {
		log.Fatalln(errQuery)
	}

	defer rows.Close() // Assurez-vous de fermer les lignes après les avoir utilisées
	var emails []string

	for rows.Next() {
		var email string
		err := rows.Scan(&email)
		if err != nil {
			log.Fatal(err)
		}
		emails = append(emails, email) // Ajoutez chaque email à la liste des emails
	}

	return emails
}

func getAllEmployees() []Employees {
	db := getDataBase()

	defer db.Close()

	rows, errQuery := db.Query("SELECT * FROM employees ORDER BY employees.lastName ASC")
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

	db := getDataBase()
	defer db.Close()

	insertStatement := "INSERT INTO employees (postId, firstName, lastName, email, password, isPresent, salary, schedule, breakTimes, dateHire, endContract) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(insertStatement)
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

func getPost(employeeID int) string {
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

func fired(employeeId string) {
	db := getDataBase()

	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM employees WHERE employeeId = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(employeeId)
	if err != nil {
		log.Fatal(err)
	}
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


func getAllDataByManager() []Employees {
	db := getDataBase()

	defer db.Close()

	rows, errQuery := db.Query("SELECT CASE WHEN employees.employeeId < 8 THEN 'Manager 1' WHEN employees.employeeId BETWEEN 10 AND 12 THEN 'Manager 2' ELSE 'Manager 3' END AS Manager, employees.* FROM employees LEFT JOIN managers AS M ON employees.employeeId = M.managerId ORDER BY Manager;")
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


func getDepartementNames(employeeID int) (string) {
	db := getDataBase()
	defer db.Close()

	query := `SELECT posts.postName, departements.serviceName FROM posts, departements  INNER JOIN employees ON posts.postId = employees.postId WHERE employees.employeeId = ?`

	var postName, serviceName string

	err := db.QueryRow(query, employeeID).Scan(&postName, &serviceName)
	if err != nil {
		log.Printf("Erreur lors de la récupération des données de département : %v", err)
		
	}

	return postName
}

func getEmployeData(id int) (Employees) {
	db := getDataBase()
	defer db.Close()

	var employee Employees 

	err := db.QueryRow("SELECT * FROM employees WHERE employeeId = ?", id).Scan(
		&employee.EmployeeId,
		&employee.PostId,
		&employee.FirstName,
		&employee.LastName,
		&employee.Email,
		&employee.Password,
		&employee.IsPresent,
		&employee.Salary,
		&employee.Schedule,
		&employee.BreackTimes,
		&employee.DateHire,
		&employee.EndContract,
	)
	if err != nil {
		log.Fatal(err)
		return Employees{}
	}

	return employee
}


func getManagers(id int) string {
    db := getDataBase()
    defer db.Close()

    query := `
        SELECT
            CASE
                WHEN employees.employeeId < 8 THEN 'Manager 1'
                WHEN employees.employeeId BETWEEN 10 AND 12 THEN 'Manager 2'
                ELSE 'Manager 3'
            END AS Manager
        FROM employees
        WHERE employeeId = ?;
    `

    var manager string

    err := db.QueryRow(query, id).Scan(&manager)
    if err != nil {
        log.Fatal(err)
    }

    return manager
}