package employee

// MemoryStore can be used for a static in-memory storage for testing purposes where an SQL library is unavailable.
type MemoryStore struct {
	employees []Employee
}

// NewMemoryStore returns a MemoryStore with an initial set of statically defined data.
func NewMemoryStore() *MemoryStore {
	var employees = []Employee{
		{Id: 11, Gender: "male"},
		{Id: 22, Gender: "male"},
		{Id: 33, Gender: "male"},
		{Id: 44, Gender: "female"},
		{Id: 55, Gender: "female"},
		{Id: 66, Gender: "female"},
	}
	return &MemoryStore{employees: employees}
}

// GetAllEmployees returns all employees currently in the store.
func (s *MemoryStore) GetAllEmployees() ([]Employee, error) {
	return s.employees, nil
}
