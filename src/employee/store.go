package employee

// Store represents the interface for any employee storage used by the Handler.
type Store interface {
	GetAllEmployees() ([]Employee, error)
}

