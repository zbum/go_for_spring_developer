package main

import "fmt"

type StudentService struct {
	studentRepository StudentRepository
}

func NewStudentService(studentRepository StudentRepository) *StudentService {
	return &StudentService{studentRepository}
}

func (r StudentService) GetStudent(id int64) *Student {
	return r.studentRepository.FindById(id)
}

func (r StudentService) RegisterStudent(student Student) error {
	if r.studentRepository.FindById(student.id) != nil {
		return fmt.Errorf("user Already Exists: %d", student.id)
	}
	r.studentRepository.Save(student)
	return nil
}

// 이하는 이 장에서 다루지 않습니다.
type Student struct {
	id   int64
	name string
}

type StudentRepository interface {
	FindById(id int64) *Student
	Save(student Student) *Student
}
