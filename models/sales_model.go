package models

import (
	"time"
)

type Sales struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    string    `gorm:"not null" json:"userId"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	Products  []Product `gorm:"many2many:sales_products;" json:"products"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `gorm:"index" json:"deletedAt,omitempty"`
}

