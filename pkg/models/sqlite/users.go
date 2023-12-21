package sqlite

import (
	"database/sql"

	"github.com/belivinge/Snippetbox/pkg/models"
	"github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

// Insert method to add a new record to the users table
func (m *UserModel) Insert(name, email, password string) error {
	// create a bcrypt hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	stmt := `INSERT INTO users (name, email, hashed_password, created) VALUES (?, ?, ?, CURRENT_TIMESTAMP)`
	// to insert the user details and hashed password into the users table
	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		if sqliteErr, ok := err.(*sqlite3.Error); ok {
			if sqliteErr.Code == 2067 {
				return models.ErrDuplicateEmail
			}
		}
	}
	return err
}

// to verify if the user exists
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get method to fetch details for a specific user
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
