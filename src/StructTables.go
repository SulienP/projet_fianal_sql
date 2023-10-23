package sql

type Departements struct {
	DepartementId int
	ServiceName   string
	StaffNumber   int
	ManagerId     int
}

type Employees struct {
	EmployeeId  int
	PostId      int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	IsPresent   bool
	Salary      int
	Schedule    string
	BreackTimes string
	DateHire    string
	EndContract string
}

type Managers struct {
	ManagerId  int
	EmployeeId int
}

type Posts struct {
	PostId        int
	DepartementId int
	PostName      string
	ManagerId     int
}
