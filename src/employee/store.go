package employee

type Store interface {
	GetAllEmployees() ([]Employee, error)
}

