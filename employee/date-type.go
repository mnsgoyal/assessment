package employee

type EmpDetails struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

type Error struct {
	Error string
}

type PaginationDetails struct {
	Previous   bool
	Next       bool
	Limit      int
	TotalPages int
	Page       int
}

type EmployeeList struct {
	PaginationDetails PaginationDetails
	EmpDetails        []EmpDetails
}
