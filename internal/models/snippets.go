package models

import (
	"database/sql"
	"time"
)


//Db schema
type Snippet struct {
	ID int
	Tile string
	Content string
	Created time.Time
	Expires time.Time
}


type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}
//Get specific
func (m * SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}
//Fetches 10 latest created snippets
func (m * SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}