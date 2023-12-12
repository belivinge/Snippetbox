package mysql

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
	return 0, nil
}

// returns specific snippet based on its id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
