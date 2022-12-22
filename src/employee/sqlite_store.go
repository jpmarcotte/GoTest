package employee

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// SQLiteStore is an SQLite specific storage for employees.
type SQLiteStore struct {
	db *sql.DB
}

// NewSQLiteStore returns an SQLiteStore with an already connected database.
func NewSQLiteStore() (*SQLiteStore, error) {
	s := SQLiteStore{}
	if err := s.connectDatabase(); err != nil {
		return nil, err
	}

	return &s, nil
}

// GetAllEmployees returns all employees currently in the store.
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

// connectDatabase is used to open the connection to the database.
func (s *SQLiteStore) connectDatabase() error {
	// TODO: Consider making this more configurable.
	// This could make the entire class more generic with a DI driver and source-name
	db, err := sql.Open("sqlite3", "./data/employees.db")
	if err != nil {
		return err
	}

	s.db = db

	return nil
}
