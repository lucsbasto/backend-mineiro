package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"type:varchar(255);unique" json:"username"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Password  string    `gorm:"type:varchar(255)" json:"-"`
	IsAdmin 	bool      `gorm:"default:false" json:"isAdmin"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `gorm:"index" json:"deletedAt,omitempty"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New().String()
	return
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	user.UpdatedAt = time.Now()
	return
}

// HashPassword gera um hash para a senha do usuário.
func (user *User) HashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}

// CheckPassword verifica se a senha fornecida corresponde ao hash armazenado.
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// Validate verifica se os campos do usuário estão válidos.
func (user *User) Validate() error {
	if user.Username == "" || user.Name == "" {
		return errors.New("nome e username são obrigatórios")
	}
	return nil
}
