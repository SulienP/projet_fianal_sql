package sql

type AllData struct {
	EmployeeId     int
	PostId         string
	FirstName      string
	LastName       string
	Email          string
	Password       string
	IsPresent      bool
	Salary         int
	Schedule       string
	BreaksTimes    string
	DateHire       string
	EndContact     string
	DepartementsId int
	PostName       string
	ManagerId      int
	ServiceId      int
	ServiceName    string
	StaffNumber    string
}

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
