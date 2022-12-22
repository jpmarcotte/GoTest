package employee

// Gender refers to an employee's Gender as a string, e.g. "female", "male", "non-binary", "other".
type Gender string // TODO: Consider restricted values?

// Employee is the data structure for storing employee information.
type Employee struct {
	Gender Gender `json:"gender"`
	Id     int    `json:"id"`
}
