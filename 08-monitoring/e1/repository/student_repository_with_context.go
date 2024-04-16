package repository

import (
	"context"
	"fmt"
	"github.com/zbum/scouter-agent-golang/scouterx/strace"
	"go_for_spring_developer/08-monitoring/e1/model"
	"gorm.io/gorm"
)

type StudentRepositoryWithContext struct{}

func NewStudentRepositoryWithContext() *StudentRepositoryWithContext {
	return &StudentRepositoryWithContext{}
}

func (s *StudentRepositoryWithContext) Save(ctx context.Context, db *gorm.DB, student *model.Student) (id uint, rowsAffected int64, err error) {
	tx := db.WithContext(ctx).Save(student)
	if tx.Error != nil {
		return 0, tx.RowsAffected, tx.Error
	}
	return student.ID, tx.RowsAffected, nil
}

func (s *StudentRepositoryWithContext) Insert(ctx context.Context, db *gorm.DB, student *model.Student) (id uint, rowsAffected int64, err error) {
	tx := db.WithContext(ctx).Create(student)
	if tx.Error != nil {
		return 0, tx.RowsAffected, tx.Error
	}
	return student.ID, tx.RowsAffected, nil
}

func (s *StudentRepositoryWithContext) FindById(ctx context.Context, db *gorm.DB, id uint) (model.Student, error) {
	step := strace.StartMethod(ctx)
	defer strace.EndMethod(ctx, step)

	var student model.Student
	tx := db.WithContext(ctx).Find(&student, id)
	if tx.Error != nil {
		return student, tx.Error
	}
	return student, nil
}

func (s *StudentRepositoryWithContext) FindAll(ctx context.Context, db *gorm.DB) (students []model.Student, err error) {
	tx := db.WithContext(ctx).Limit(100).Find(&students)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return students, nil
}

func (s *StudentRepositoryWithContext) UpdateNameById(ctx context.Context, db *gorm.DB, id uint, newName string) (int64, error) {
	tx := db.WithContext(ctx).Model(&model.Student{}).Where("ID = ?", id).Update("Name", newName)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (s *StudentRepositoryWithContext) UpdateNames(ctx context.Context, db *gorm.DB, newName string) (int64, error) {
	tx := db.WithContext(ctx).Model(&model.Student{}).Where("1=1").Updates(model.Student{Name: newName})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (s *StudentRepositoryWithContext) DeleteById(ctx context.Context, db *gorm.DB, id uint) (int64, error) {
	tx := db.WithContext(ctx).Delete(&model.Student{}, id)
	if tx.Error != nil {
		fmt.Println(db.Error)
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (s *StudentRepositoryWithContext) DeleteAll(ctx context.Context, db *gorm.DB) (int64, error) {
	tx := db.WithContext(ctx).Where("1 = 1").Delete(&model.Student{})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}
