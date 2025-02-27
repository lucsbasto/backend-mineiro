package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)
type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"type:varchar(255);unique" json:"username"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `gorm:"index" json:"deletedAt,omitempty"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
    user.ID = uuid.New().String()
    return
}