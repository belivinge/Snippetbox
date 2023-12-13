package sqlite

import (
	"database/sql"

	"github.com/belivinge/Snippetbox/pkg/models"
)

// snippet model type which wraps from mysql sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// insert a new snippet into db
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	// sql statement we want to execute
	stmt := `INSERT INTO snippets (title, content, created, expires) VALUES (?, ?, CURRENT_TIMESTAMP, ?)`

	// this method returns a sql.Result object
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	// to get the id of our inserted record in the table
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// the id has type int64, we turn it to int
	return int(id), err
}

// returns specific snippet based on its id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
