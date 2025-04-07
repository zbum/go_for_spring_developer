package repository

import (
	"fmt"
	"go_for_spring_developer/08-database/01-common/model"
	"gorm.io/gorm"
)

type StudentRepository struct{}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{}
}

func (s *StudentRepository) Save(db *gorm.DB, student *model.Student) (id uint, rowsAffected int64, err error) {
	tx := db.Save(student)
	if tx.Error != nil {
		return 0, tx.RowsAffected, tx.Error
	}
	return student.ID, tx.RowsAffected, nil
}

func (s *StudentRepository) Insert(db *gorm.DB, student *model.Student) (id uint, rowsAffected int64, err error) {
	tx := db.Create(student)
	if tx.Error != nil {
		return 0, tx.RowsAffected, tx.Error
	}
	return student.ID, tx.RowsAffected, nil
}

func (s *StudentRepository) FindById(db *gorm.DB, id uint) (model.Student, error) {
	var student model.Student
	tx := db.Find(&student, id)
	if tx.Error != nil {
		return student, tx.Error
	}
	return student, nil
}

func (s *StudentRepository) FindAll(db *gorm.DB) (students []model.Student, err error) {
	tx := db.Model(&model.Student{}).Preload("Scores").Find(&students)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return students, nil
}

func (s *StudentRepository) UpdateNameById(db *gorm.DB, id uint, newName string) (int64, error) {
	tx := db.Model(&model.Student{}).Where("ID = ?", id).Update("Name", newName)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (s *StudentRepository) UpdateNames(db *gorm.DB, newName string) (int64, error) {
	tx := db.Model(&model.Student{}).Where("1=1").Updates(model.Student{Name: newName})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (s *StudentRepository) DeleteById(db *gorm.DB, id uint) (int64, error) {
	tx := db.Delete(&model.Student{}, id)
	if tx.Error != nil {
		fmt.Println(db.Error)
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}

func (s *StudentRepository) DeleteAll(db *gorm.DB) (int64, error) {
	tx := db.Where("1 = 1").Delete(&model.Student{})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}
