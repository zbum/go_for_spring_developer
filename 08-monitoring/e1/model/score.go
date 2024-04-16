package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Score struct {
	gorm.Model
	Score     uint
	StudentID uint
}

func (Score) TableName() string {
	return "Scores"
}

func (s Score) String() string {
	return fmt.Sprintf("ID: %d, Score: %d, DeletedAt: %v\n", s.ID, s.Score, s.DeletedAt)
}
