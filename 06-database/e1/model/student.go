package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string
}

func (Student) tableName() string {
	return "Students"
}

func (s Student) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, DeletedAt: %v\n", s.ID, s.Name, s.DeletedAt)
}
