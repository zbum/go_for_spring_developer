package model

import (
	"fmt"
)

type Student struct {
	ID     uint `gorm:"primarykey"`
	Name   string
	Age    uint
	Scores []Score `gorm:"foreignKey:StudentID;references:ID"`
}

func (Student) TableName() string {
	return "Students"
}

func (s Student) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, Age: %v, Scores: %v \n", s.ID, s.Name, s.Age, s.Scores)
}
