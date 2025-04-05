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
}

func insertSampleData(db *gorm.DB) []uint {
	// loglevel 을 잠시 낮춥니다.
	//db.Logger.LogMode(logger.Silent)

	var ids []uint

	// 생성
	id, _ := crud.Insert(db, &model.Student{
		Name: "Manty1",
		Age:  15,
		Scores: []model.Score{
			{Score: 10},
			{Score: 11},
		},
	})
	ids = append(ids, id)

	id, _ = crud.Insert(db, &model.Student{Name: "Manty2", Age: 14})
	ids = append(ids, id)

	id, _ = crud.Insert(db, &model.Student{Name: "Manty3", Age: 13})
	ids = append(ids, id)

	id, _ = crud.Insert(db, &model.Student{Name: "Manty4", Age: 12})
	ids = append(ids, id)

	id, _ = crud.Insert(db, &model.Student{Name: "Manty5", Age: 11})
	ids = append(ids, id)

	id, _ = crud.Insert(db, &model.Student{Name: "Manty6", Age: 10})
	ids = append(ids, id)

	id, _ = crud.Insert(db, &model.Student{Name: "Manty7", Age: 9})
	ids = append(ids, id)

	id, _ = crud.Insert(db, &model.Student{Name: "Manty8", Age: 8})
	ids = append(ids, id)

	id, _ = crud.Insert(db, &model.Student{Name: "Manty9", Age: 7})
	ids = append(ids, id)

	id, _ = crud.Insert(db, &model.Student{Name: "Manty10", Age: 6})
	ids = append(ids, id)
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
		Where("name = ?", "Manty1").
		Find(&selectedStudents)
	fmt.Println("[1] Eager Loaded Student : \n", selectedStudents)

	// Limit 를 취소할때.
	db.Limit(3).Order("name asc").Order("age desc").Limit(-1).Find(&selectedStudents)
	fmt.Println("[2] ConditionedSortedLimitedSelected Student : \n", selectedStudents)

	// Limit 와 offset 을 함께 사용
	db.Limit(3).Offset(5).Order("name asc").Order("age desc").Find(&selectedStudents)
	fmt.Println("[3] ConditionedSortedLimitedSelected Student : \n", selectedStudents)

	deleteSampleData(db)
}
