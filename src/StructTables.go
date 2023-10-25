package sql

type Departements struct {
	DepartementIdDepartements int
	ServiceName   string
	StaffNumberDepartements   int
	ManagerIdDepartements     int
	DepartementId int
	ServiceNameDepartement   string
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
	IsPresent   string
	Salary      int
	Schedule    string
	BreackTimes string
	DateHire    string
	EndContract string
}

type Managers struct {
	ManagerIdFromManager  int
	EmployeeIdFromManage int
	ManagerId  int
	EmployeeId int
}

type Posts struct {
	PostIdFromPosts        int
	PostNameFromPost     string
	ManagerIdFromPosts    int
	DepartementId int
	PostName      string
	ManagerId     int
}
