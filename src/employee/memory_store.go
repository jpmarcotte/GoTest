package employee

type MemoryStore struct {
	employees []Employee
}

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

func (s *MemoryStore) GetAllEmployees() ([]Employee, error) {
	return s.employees, nil
}
