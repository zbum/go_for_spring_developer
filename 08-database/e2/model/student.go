package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string
	Scores []Score
}

func (Student) TableName() string {
	return "Students"
}

func (s Student) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, DeletedAt: %v, Scores: %v \n", s.ID, s.Name, s.DeletedAt, s.Scores)
}
