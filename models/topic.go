package models

import (
	"database/sql"
	"time"
)

type TopicRequest struct {
	Name string `json:"name"`
}

type Topic struct {
	ID        int          `db:"id"`
	Name      string       `db:"name"`
	Slug      string       `db:"slug"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type TopicFilterParameters struct {
	CurrentPage int
	PerPage     int
	Offset      int
}
