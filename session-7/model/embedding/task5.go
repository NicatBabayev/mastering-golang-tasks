package embedding

type Person struct {
	FirstName string
	LastName  string
}

type Employee struct {
	EmployeeID int
	Position   string
	Person
}
