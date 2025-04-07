package model

import (
	"fmt"
)

type Score struct {
	ID        uint
	Score     uint
	StudentID uint
}

func (Score) TableName() string {
	return "Scores"
}

func (s Score) String() string {
	return fmt.Sprintf("ID: %d, Score: %d", s.ID, s.Score)
}
