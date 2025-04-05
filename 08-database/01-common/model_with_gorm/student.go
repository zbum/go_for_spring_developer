package model_with_gorm

import (
	"fmt"
	"gorm.io/gorm"
)

type StudentWithDeletedAt struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Scores    []ScoreWithGormModel `gorm:"foreignKey:StudentID;references:ID"`
	DeletedAt gorm.DeletedAt       `gorm:"index"`
}

func (StudentWithDeletedAt) TableName() string {
	return "StudentsWithGormModel"
}

func (s StudentWithDeletedAt) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, DeletedAt: %v, Scores: %v \n", s.ID, s.Name, s.DeletedAt, s.Scores)
}

type StudentWithGormModel struct {
	gorm.Model
	Name   string
	Scores []ScoreWithGormModel `gorm:"foreignKey:StudentID;references:ID"`
}

func (StudentWithGormModel) TableName() string {
	return "StudentsWithGormModel"
}

func (s StudentWithGormModel) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, DeletedAt: %v, Scores: %v \n", s.ID, s.Name, s.DeletedAt, s.Scores)
}
