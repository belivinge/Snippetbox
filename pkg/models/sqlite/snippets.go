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
	var stmt string
	if expires == "1" {
		stmt = `INSERT INTO snippets (title, content, created, expires)
		VALUES(?, ?, CURRENT_TIMESTAMP, DATE(CURRENT_TIMESTAMP, '+1 DAY'))`
	} else if expires == "7" {
		stmt = `INSERT INTO snippets (title, content, created, expires)
		VALUES(?, ?, CURRENT_TIMESTAMP, DATE(CURRENT_TIMESTAMP, '+7 DAYS'))`
	} else {
		stmt = `INSERT INTO snippets (title, content, created, expires)
		VALUES(?, ?, CURRENT_TIMESTAMP, DATE(CURRENT_TIMESTAMP, '+1 YEAR'))`
	}
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
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE id = ?`
	// this does not work in sqlite
	// stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > CURRENT_TIMESTAMP AND id = ?`
	// to execute SQL statement
	row := m.DB.QueryRow(stmt, id)

	// initializing a pointer to Snippet struct
	s := &models.Snippet{}
	// err := m.DB.QueryRow("SELECT * FROM snippets WHERE id = ?", id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}

// 10 most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	// stmt we want to execute
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > CURRENT_TIMESTAMP ORDER BY created DESC LIMIT 10`
	// Query() method to execute sql stmt
	// returns sql.Rows
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	// ensure that sql.rows is closed before func closes
	defer rows.Close()

	// initialize slice to hold models
	snippets := []*models.Snippet{}

	// rows.Next() to iterate the rows in the resultset, prepares the first row to be acted
	for rows.Next() {
		// pointer to snippet struct
		s := &models.Snippet{}
		// rows.scan to copy the values from each field to the new snippet
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)

	}
	// don't assume that iteration was cmpleted
	if err = rows.Err(); err != nil {
		return nil, err
	}
	// snippets slice
	return snippets, nil
}
