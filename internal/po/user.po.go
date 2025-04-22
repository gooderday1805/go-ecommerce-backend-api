package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"column:uuid;type:varchar(36);not null;uniqueIndex"`
	Name     string    `gorm:"column:name"`
	IsActive bool      `gorm:"column:is_active; type:boolean"`
	Roles    []Role    `gorm:"many2many:go_user_roles;"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
