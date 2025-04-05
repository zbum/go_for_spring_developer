package main

import (
	"fmt"
	"go_for_spring_developer/08-database/01-common/crud"
	"go_for_spring_developer/08-database/01-common/db"
	"go_for_spring_developer/08-database/01-common/model_with_gorm"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db := db.InitGorm()
	queryWithGormModel(db)
}

func insertSampleData(db *gorm.DB) []uint {
	// loglevel 을 잠시 낮춥니다.
	db.Logger.LogMode(logger.Silent)

	var ids []uint

	// 생성
	id, _ := crud.InsertWithGormModel(db, &model_with_gorm.StudentWithGormModel{Name: "Manty1"})
	ids = append(ids, id)

	id, _ = crud.InsertWithGormModel(db, &model_with_gorm.StudentWithGormModel{Name: "Manty2"})
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
	fmt.Println("[6] Deleted All")

	// loglevel 복구
	db.Logger.LogMode(logger.Info)
}

func queryWithGormModel(db *gorm.DB) {
	fmt.Println("\n\n#### START queryWithGormModel ####")

	var ids = insertSampleData(db)

	// 생성된 QUERY 를 살펴 보자. ( DeletedAt 을 사용한 경우)
	var selectedStudent model_with_gorm.StudentWithDeletedAt
	lastId := ids[len(ids)-1]
	db.First(&selectedStudent, lastId)
	fmt.Println("[2] DeletedAt using Selected Student : ", selectedStudent)

	// IN 조건으로 조회
	var inSelectedStudents []model_with_gorm.StudentWithGormModel
	db.Find(&inSelectedStudents, ids)
	fmt.Println("[3] GormModel Student : ", inSelectedStudents)

	deleteSampleData(db)
}
