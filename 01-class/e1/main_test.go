package main

import "testing"

func TestUserRepository_AddUser_FindById(t *testing.T) {

	student := &Student{1, "dummy"}

	g := NewStudentService(&DummyStudentRepository{})

	t.Run("GetStudent", func(t *testing.T) {
		if got := g.GetStudent(1); *got != *student {
			t.Errorf("FindById() = %v, want %v", got, student)
		}
	})
}

func TestUserRepository_Find_By_ExistId(t *testing.T) {

	g := NewStudentService(&DummyStudentRepository{})

	t.Run("RegisterStudent", func(t *testing.T) {

		if got := g.RegisterStudent(*student1); got == nil {
			t.Errorf("RegisterStudent() = %v, want %v", got, nil)
		}
	})
}

// 이하는 이 장에서 다루지 않습니다.

var student1 = &Student{1, "dummy"}

type DummyStudentRepository struct {
}

func (r DummyStudentRepository) FindById(id int64) *Student {
	return student1
}
func (r DummyStudentRepository) Save(student Student) *Student {
	return &student
}
