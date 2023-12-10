package todo_model

import (
	"time"
)

type Todo struct {
	ID         int64
	Title      string
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeleteFlag bool
}
