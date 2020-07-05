package models

import (
	"database/sql"
	"time"
)

type NewsRequest struct {
	TopicID     int    `json:"topic_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	IsPublish   bool   `json:"is_publish"`
}

type News struct {
	ID          int          `db:"id"`
	TopicID     int          `db:"topic_id"`
	Title       string       `db:"title"`
	Slug        string       `db:"slug"`
	Description string       `db:"description"`
	Tags        string       `db:"tags"`
	IsPublish   bool         `db:"is_publish"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
	Topic       Topic        `gorm:"foreignkey:TopicID;association_save_reference:false;association_autoupdate:false"`
}

type NewsFilterParams struct {
	Title       string
	Topic       []int
	Tags        []string
	Published   string
	Deleted     string
	CurrentPage int
	PerPage     int
	Offset      int
}
