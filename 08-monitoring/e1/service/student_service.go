package service

import (
	"context"
	"github.com/zbum/scouter-agent-golang/scouterx/strace"
	"go_for_spring_developer/08-monitoring/e1/configuration/database"
	"go_for_spring_developer/08-monitoring/e1/model"
	"go_for_spring_developer/08-monitoring/e1/repository"
	"gorm.io/gorm"
)

type StudentService struct {
	datasource        *database.Datasource
	studentRepository *repository.StudentRepositoryWithContext
}

func NewStudentService(datasource *database.Datasource, studentRepository *repository.StudentRepositoryWithContext) *StudentService {
	return &StudentService{datasource: datasource, studentRepository: studentRepository}
}

func (s *StudentService) GetStudent(ctx context.Context, id uint) (model.Student, error) {
	step := strace.StartMethod(ctx)
	defer strace.EndMethod(ctx, step)

	db := s.datasource.GetDB()
	return s.studentRepository.FindById(ctx, db, id)
}

func (s *StudentService) RegisterStudent(ctx context.Context, student *model.Student) (uint, error) {
	var studentId uint
	db := s.datasource.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {

		id, _, err := s.studentRepository.Save(ctx, tx, student)
		studentId = id
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return studentId, nil
}
