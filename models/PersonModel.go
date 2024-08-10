package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {

	gorm.Model
	ID uint
	Name string
	Address string
	Phone uint 
	CreatedAt time.Time
	UpdatedAt time.Time

}

func (Person) TableName() string {

	return "persons"

}