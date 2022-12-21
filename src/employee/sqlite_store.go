package employee

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore() (*SQLiteStore, error) {
	s := SQLiteStore{}
	if err := s.connectDatabase(); err != nil {
		return nil, err
	}

	return &s, nil
}

func (s *SQLiteStore) GetAllEmployees() ([]Employee, error) {
	rows, err := s.db.Query("SELECT id, gender FROM employees")
	if err != nil {
		return nil, err
	}

	//goland:noinspection GoUnhandledErrorResult nothing useful to do with this error
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

func (s *SQLiteStore) connectDatabase() error {
	db, err := sql.Open("sqlite3", "./data/employees.db")
	if err != nil {
		return err
	}

	s.db = db

	return nil
}
