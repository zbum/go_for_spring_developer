package repository

import (
	"go_for_spring_developer/08-database/e1/model"
	"gorm.io/gorm"
)

type ScoreRepository struct{}

func NewScoreRepository() *ScoreRepository {
	return &ScoreRepository{}
}

func (s ScoreRepository) Save(db *gorm.DB, score *model.Score) (id uint, rowsAffected int64, err error) {
	tx := db.Save(score)
	if tx.Error != nil {
		return 0, tx.RowsAffected, tx.Error
	}
	return score.ID, tx.RowsAffected, nil
}
