package table

import (
	"github.com/google/uuid"
)

// UpperCamelCaseで書き、Table,columnはsnake_caseの複数形でmigrationされる
type User struct {
	// gorm.Modelをつけると、incrementalなidとCreatedAtとUpdatedAtとDeletedAtが作られる
	Id         uuid.UUID `gorm:"size:36;index:idx_user_id"`
	Email      string    `gorm:"size:255;index:idx_email,unique"`
	Age        int       `gorm:"check:age >= 0"`
	DeleteFlag bool
	Todos      Todo `gorm:"constraint:OnDelete:CASCADE;"`
}

type Todo struct {
	Id uuid.UUID `gorm:"size:36;index:idx_todo_id"`

	Title      string `gorm:"size:255;"`
	Content    string `gorm:"size:255;"`
	DeleteFlag bool
	UserId     uuid.UUID
}
