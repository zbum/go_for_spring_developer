package crud

import (
	"fmt"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string
}

func (Student) tableName() string {
	return "Students"
}

func (s Student) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, DeletedAt: %v\n", s.ID, s.Name, s.DeletedAt)
}

func Save(db *gorm.DB, student *Student) (id uint, rowsAffected int64) {
	tx := db.Save(student)
	if tx.Error != nil {
		fmt.Println(db.Error)
	}
	return student.ID, tx.RowsAffected
}

func Insert(db *gorm.DB, student *Student) (id uint, rowsAffected int64) {
	tx := db.Create(student)
	if tx.Error != nil {
		fmt.Println(db.Error)
	}
	return student.ID, tx.RowsAffected
}

func FindById(db *gorm.DB, id uint) Student {
	var student Student
	tx := db.Find(&student, id)
	if tx.Error != nil {
		fmt.Println(db.Error)
		return student
	}
	return student
}

func FindAll(db *gorm.DB) (students []Student) {
	tx := db.Limit(100).Find(&students)
	if tx.Error != nil {
		fmt.Println(db.Error)
	}
	return students
}

func UpdateNameById(db *gorm.DB, id uint, newName string) int64 {
	tx := db.Model(&Student{}).Where("ID = ?", id).Update("Name", newName)
	if tx.Error != nil {
		fmt.Println(db.Error)
		return 0
	}
	return tx.RowsAffected
}

func UpdateNames(db *gorm.DB, newName string) int64 {
	tx := db.Model(&Student{}).Where("1=1").Updates(Student{Name: newName})
	if tx.Error != nil {
		fmt.Println(db.Error)
		return 0
	}
	return tx.RowsAffected
}

func DeleteById(db *gorm.DB, id uint) int64 {
	tx := db.Delete(&Student{}, id)
	if tx.Error != nil {
		fmt.Println(db.Error)
		return 0
	}
	return tx.RowsAffected
}

func DeleteAll(db *gorm.DB) int64 {
	tx := db.Where("1 = 1").Delete(&Student{})
	if tx.Error != nil {
		fmt.Println(db.Error)
		return 0
	}
	return tx.RowsAffected
}
