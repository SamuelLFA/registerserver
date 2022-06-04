package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Person struct {
	Id       string `gorm:"primaryKey"`
	Name     string `gorm:"notNull"`
	Document string `gorm:"notNull;unique"`
	Email    string `gorm:"notNull;unique"`
}

func (person *Person) BeforeCreate(tx *gorm.DB) (err error) {
	person.Id = uuid.NewString()
	return
}
