package table

import (
	"gorm.io/gorm"
)

type User struct {
	// gorm.Modelをつけると、idとCreatedAtとUpdatedAtとDeletedAtが作られる
	gorm.Model

	Name       string `gorm:"size:255;index:idx_name"`
	Email      string `gorm:"size:255;index:idx_email,unique"`
	Age        int
	DeleteFlag bool
	Todos      []Todo
}

type Todo struct {
	gorm.Model

	Title      string
	Content    string
	DeleteFlag bool
	UserID     uint
}
