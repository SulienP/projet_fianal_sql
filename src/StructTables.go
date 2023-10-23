package sql

type Departements struct {
	departementId int
	serviceName   string
	staffNumber   int
	managerId     int
}

type Employees struct {
	employeeId  int
	postId      int
	firstName   string
	lastName    string
	email       string
	password    string
	isPresent   bool
	salary      int
	schedule    string
	breackTimes string
	dateHire    string
	endContract string
}

type Managers struct {
	managerId  int
	employeeId int
}

type Posts struct {
	postId        int
	departementId int
	postName      string
	managerId     int
}
