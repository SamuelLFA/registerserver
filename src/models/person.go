package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Person struct {
	Id       string `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"notNull"`
	Document string `json:"document" gorm:"notNull"`
	Email    string `json:"email" gorm:"notNull"`
}

func (person *Person) BeforeCreate(tx *gorm.DB) (err error) {
	person.Id = uuid.NewString()
	return
}
