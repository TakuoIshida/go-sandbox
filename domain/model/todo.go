package todomodel

import (
	"log"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id uuid.UUID

	Title      string
	Content    string
	DeleteFlag bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserId     uuid.UUID
}

func New(title string, content string, userId uuid.UUID) Todo {
	uuid, err := uuid.NewV7()
	if err != nil {
		log.Fatal(err.Error())
	}
	// validate
	return Todo{
		Id:         uuid,
		Title:      title,
		Content:    content,
		UserId:     userId,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		DeleteFlag: false,
	}
}

func Restore(t *Todo) *Todo {

	// validate
	return &Todo{
		Id:         t.Id,
		Title:      t.Title,
		Content:    t.Content,
		UserId:     t.UserId,
		CreatedAt:  t.CreatedAt,
		UpdatedAt:  t.UpdatedAt,
		DeleteFlag: t.DeleteFlag,
	}
}

// domain rule
// func validate
