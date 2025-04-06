package crud

import (
	"fmt"
	"go_for_spring_developer/08-database/01-common/model"
	"go_for_spring_developer/08-database/01-common/model_with_gorm"
	"gorm.io/gorm"
)

func Save(db *gorm.DB, student *model.Student) (id uint, rowsAffected int64) {
	tx := db.Save(student)
	if tx.Error != nil {
		fmt.Println(db.Error)
	}
	return student.ID, tx.RowsAffected
}

func Insert(db *gorm.DB, student *model.Student) (id uint, rowsAffected int64) {
	tx := db.Create(student)
	if tx.Error != nil {
		fmt.Println(db.Error)
	}
	return student.ID, tx.RowsAffected
}

func InsertWithGormModel(db *gorm.DB, student *model_with_gorm.StudentWithGormModel) (id uint, rowsAffected int64) {
	tx := db.Create(student)
	if tx.Error != nil {
		fmt.Println(db.Error)
	}
	return student.ID, tx.RowsAffected
}

func Inserts(db *gorm.DB, students []*model.Student) (rowsAffected int64) {
	tx := db.Create(students)
	if tx.Error != nil {
		fmt.Println(db.Error)
	}
	return tx.RowsAffected
}

func FindById(db *gorm.DB, id uint) model.Student {
	var student model.Student
	tx := db.Find(&student, id)
	if tx.Error != nil {
		fmt.Println(db.Error)
		return student
	}
	return student
}

func FindAll(db *gorm.DB) (students []model.Student) {
	tx := db.Limit(100).Find(&students)
	if tx.Error != nil {
		fmt.Println(db.Error)
	}
	return students
}

func UpdateNameById(db *gorm.DB, id uint, newName string) int64 {
	tx := db.Model(&model.Student{}).Where("ID = ?", id).Update("Name", newName)
	if tx.Error != nil {
		fmt.Println(db.Error)
		return 0
	}
	return tx.RowsAffected
}

func UpdateNames(db *gorm.DB, newName string) int64 {
	tx := db.Model(&model.Student{}).Where("1=1").Updates(model.Student{Name: newName})
	if tx.Error != nil {
		fmt.Println(db.Error)
		return 0
	}
	return tx.RowsAffected
}

func DeleteById(db *gorm.DB, id uint) int64 {
	tx := db.Delete(&model.Student{}, id)
	if tx.Error != nil {
		fmt.Println(db.Error)
		return 0
	}
	return tx.RowsAffected
}

func DeleteAll(db *gorm.DB) {
	db.Transaction(func(tx *gorm.DB) error {

		tx.Where("1 = 1").Delete(&model.Score{})
		if tx.Error != nil {
			fmt.Println(db.Error)
			return tx.Error
		}

		tx.Where("1 = 1").Delete(&model.Student{})
		if tx.Error != nil {
			fmt.Println(db.Error)
			return tx.Error
		}

		tx.Where("1 = 1").Delete(&model_with_gorm.ScoreWithGormModel{})
		if tx.Error != nil {
			fmt.Println(db.Error)
			return tx.Error
		}

		tx.Where("1 = 1").Delete(&model_with_gorm.StudentWithGormModel{})
		if tx.Error != nil {
			fmt.Println(db.Error)
			return tx.Error
		}
		return nil
	})

}
