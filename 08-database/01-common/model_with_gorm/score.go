package model_with_gorm

import (
	"fmt"
	"gorm.io/gorm"
)

type ScoreWithGormModel struct {
	gorm.Model
	Score     uint
	StudentID uint
}

func (ScoreWithGormModel) TableName() string {
	return "ScoresWithGormModel"
}

func (s ScoreWithGormModel) String() string {
	return fmt.Sprintf("ID: %d, Score: %d, DeletedAt: %v\n", s.ID, s.Score, s.DeletedAt)
}
