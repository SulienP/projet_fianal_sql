package sql

import "log"

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
