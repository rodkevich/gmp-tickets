package ticket

import "github.com/google/uuid"

type Tag struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Searchable bool      `json:"searchable"`
}

type Tags struct {
	Count int   `json:"count"`
	List  []Tag `json:"list"`
}
