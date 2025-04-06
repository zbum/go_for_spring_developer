package main

import (
	"fmt"
	"go_for_spring_developer/08-database/01-common/crud"
	"go_for_spring_developer/08-database/01-common/db"
	"go_for_spring_developer/08-database/01-common/model"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db := db.InitGorm()
	queryWithJoin(db)
	queryWithUserSpecifiedJoinCondition(db)
}

func insertSampleData(db *gorm.DB) []uint {
	// loglevel 을 잠시 낮춥니다.
	//db.Logger.LogMode(logger.Silent)

	var ids []uint

	// 생성
	students := []*model.Student{
		{
			Name: "Manty01",
			Age:  14,
			Scores: []model.Score{
				{Score: 10},
				{Score: 11},
			},
		},
		{
			Name: "Manty02",
			Age:  15,
			Scores: []model.Score{
				{Score: 10},
				{Score: 11},
			},
		},
		{
			Name: "Manty03",
			Age:  14,
			Scores: []model.Score{
				{Score: 10},
				{Score: 11},
			},
		},
		{
			Name: "Manty04",
			Age:  15,
			Scores: []model.Score{
				{Score: 10},
				{Score: 11},
			},
		},
		{
			Name: "Manty05",
			Age:  13,
			Scores: []model.Score{
				{Score: 10},
				{Score: 11},
			},
		},
		{
			Name: "Manty06",
			Age:  15,
			Scores: []model.Score{
				{Score: 10},
				{Score: 11},
			},
		},
		{
			Name: "Manty07",
			Age:  14,
			Scores: []model.Score{
				{Score: 10},
				{Score: 11},
			},
		},
		{
			Name: "Manty08",
			Age:  15,
			Scores: []model.Score{
				{Score: 10},
				{Score: 11},
			},
		},
		{
			Name: "Manty09",
			Age:  16,
			Scores: []model.Score{
				{Score: 10},
				{Score: 11},
			},
		},
		{
			Name: "Manty10",
			Age:  15,
			Scores: []model.Score{
				{Score: 10},
				{Score: 11},
			},
		},
	}

	crud.Inserts(db, students)

	for _, student := range students {
		ids = append(ids, student.ID)
	}

	fmt.Println("[1] Inserted IDs, Count : ", ids, len(ids))

	// loglevel 복구
	db.Logger.LogMode(logger.Info)

	return ids
}

func deleteSampleData(db *gorm.DB) {
	// loglevel 을 잠시 낮춥니다.
	db.Logger.LogMode(logger.Silent)

	// 모든 테스트 데이터 삭제
	crud.DeleteAll(db)
	fmt.Println("[6] Deleted All : ")

	// loglevel 복구
	db.Logger.LogMode(logger.Info)
}

func queryWithJoin(db *gorm.DB) {
	fmt.Println("\n\n#### START queryWithJoin ####")

	insertSampleData(db)

	// eager loading
	var selectedStudents []model.Student
	db.Model(&model.Student{}).
		Preload("Scores").
		Where("age = ?", 15).
		Find(&selectedStudents)
	fmt.Println("[1] Eager Loaded Student : \n", selectedStudents)

	deleteSampleData(db)
}

type ResultOfStudentWithScores struct {
	Id    uint
	Name  string
	Age   uint
	Score uint
}

func (r ResultOfStudentWithScores) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, Age: %v, Score: %v \n", r.Id, r.Name, r.Age, r.Score)
}

func queryWithUserSpecifiedJoinCondition(db *gorm.DB) {
	fmt.Println("\n\n#### START queryWithUserSpecifiedJoinCondition ####")

	insertSampleData(db)

	var resultOfStudentWithScores []ResultOfStudentWithScores
	db.Model(&model.Score{}).
		Select("Students.id, Students.age, Scores.score").
		Joins("left join Students on Scores.student_id = Students.id").
		Where("Students.age = ?", 15).
		Scan(&resultOfStudentWithScores)
	fmt.Println("[2] User Specified Join Students : \n", resultOfStudentWithScores)

	deleteSampleData(db)
}
