package employee

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Included to support sqlite storage drivers
)

type SQLStore struct {
	db *sql.DB
}

func NewSQLStore(driver, dataSource string) (*SQLStore, error) {
	s := SQLStore{}
	if err := s.connectDatabase(driver, dataSource); err != nil {
		return nil, err
	}

	return &s, nil
}

func (s *SQLStore) GetAllEmployees() ([]Employee, error) {
	rows, err := s.db.Query("SELECT id, gender FROM employees")
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer rows.Close()

	employees := make([]Employee, 0)

	for rows.Next() {
		employee := Employee{}
		err = rows.Scan(&employee.Id, &employee.Gender)
		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

func (s *SQLStore) connectDatabase(driver, dataSource string) error {
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		return err
	}

	s.db = db

	return nil
}
