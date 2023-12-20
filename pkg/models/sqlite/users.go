package sqlite

import (
	"database/sql"

	"github.com/belivinge/Snippetbox/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

// Insert method to add a new record to the users table
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// to verify if the user exists
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get method to fetch details for a specific user
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
