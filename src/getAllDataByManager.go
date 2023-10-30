// SELECT
//     CASE
//         WHEN Employee.employeeId < 8 THEN 'Manager 1'
//         WHEN Employee.employeeId BETWEEN 10 AND 12 THEN 'Manager 2'
//         ELSE 'Manager 3'
//     END AS Manager,
//     Employee.FirstName,
//     Employee.LastName
// FROM employees AS Employee
// LEFT JOIN managers AS m ON Employee.employeeId = m.managerId
// ORDER BY Manager;
package sql

import (
	"log"
)

func getAllDataByManager() ([]Employees) {
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