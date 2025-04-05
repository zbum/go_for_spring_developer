package main

import (
	"fmt"
	"go_for_spring_developer/08-database/01-common/crud"
	"go_for_spring_developer/08-database/01-common/db"
	"go_for_spring_developer/08-database/01-common/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"strconv"
)

func main() {
	db := db.InitGorm()
	querySingle(db)
	queryByPrimaryKey(db)
	queryByCondition(db)
	queryWithSort(db)
	queryWithLimitAndOffset(db)
}

func insertSampleData(db *gorm.DB) []uint {
	// loglevel 을 잠시 낮춥니다.
	db.Logger.LogMode(logger.Silent)

	var ids []uint

	// 생성
	id, _ := crud.Insert(db, &model.Student{Name: "Manty1", Age: 15})
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
	fmt.Println("[6] Deleted All")

	// loglevel 복구
	db.Logger.LogMode(logger.Info)
}

func querySingle(db *gorm.DB) {

	fmt.Println("\n\n#### START querySingle ####")

	insertSampleData(db)

	// pk 로 정렬한 첫번째 데이터 조회
	var selectedStudent model.Student
	tx := db.First(&selectedStudent)
	fmt.Printf("[3] SingleSelected Student : ROWs: %d, ERROR: %v\n", tx.RowsAffected, tx.Error)
	fmt.Println("[3] SingleSelected Student : ", selectedStudent)

	// 아무 데이터나 1개 조회 (정렬 없음)
	db.Take(&selectedStudent)
	fmt.Println("[3] SingleSelected Student : ", selectedStudent)

	deleteSampleData(db)
}

func queryByPrimaryKey(db *gorm.DB) {
	fmt.Println("\n\n#### START queryByPrimaryKey ####")

	var ids = insertSampleData(db)

	// 마지막 insert한 데이터만 조회
	var selectedStudent model.Student
	lastId := ids[len(ids)-1]
	db.First(&selectedStudent, lastId)
	fmt.Println("[2] PrimaryKey Selected Student : ", selectedStudent)

	// 마지막 insert한 데이터만 조회
	// 문자열로 조회해도 동작합니다.
	var selectedStudentByString model.Student
	db.Take(&selectedStudentByString, strconv.FormatUint(uint64(lastId), 10))
	fmt.Println("[3] PrimaryKey Selected Student : ", selectedStudentByString)

	// IN 조건으로 조회
	var inSelectedStudents []model.Student
	db.Find(&inSelectedStudents, ids)
	fmt.Println("[4] In Selected Student : ", inSelectedStudents)

	// 목적지 변수를 활용
	var selectedStudentWithPk model.Student
	selectedStudentWithPk.ID = ids[0]
	db.First(&selectedStudentWithPk)
	fmt.Println("[5] 목적지 변수를 활용 : ", selectedStudentWithPk)

	deleteSampleData(db)
}

func queryByCondition(db *gorm.DB) {
	fmt.Println("\n\n#### START queryByCondition ####")

	ids := insertSampleData(db)

	// Manty2 인 데이터만 조회
	var selectedStudents []model.Student
	db.Where("name = ?", "Manty2").Find(&selectedStudents)
	fmt.Println("[1] ConditionedSelected Student : ", selectedStudents)

	// Manty2 가 아닌 데이터만 조회
	db.Where("name <> ?", "Manty2").Find(&selectedStudents)
	fmt.Println("[2] ConditionedSelected Student : ", selectedStudents)

	// Manty1 과 Manty2 를 IN 절로 조회
	db.Where("name IN ?", []string{"Manty1", "Manty2"}).Find(&selectedStudents)
	fmt.Println("[3] ConditionedSelected Student : ", selectedStudents)

	// LIKE 사용
	db.Where("name LIKE ?", "Man%").Find(&selectedStudents)
	fmt.Println("[4] ConditionedSelected Student : ", selectedStudents)

	// AND 사용
	db.Where("name = ? AND age = ?", "Manty1", 15).Find(&selectedStudents)
	fmt.Println("[5] ConditionedSelected Student : ", selectedStudents)

	// 구조체 사용
	db.Where(&model.Student{Name: "Manty2", Age: 14}).Find(&selectedStudents)
	fmt.Println("[5] Struct ConditionedSelected Student : ", selectedStudents)

	// Map 사용
	mapCondition := map[string]interface{}{"name": "Manty1"}
	db.Where(mapCondition).Find(&selectedStudents)
	fmt.Println("[6] Map ConditionedSelected Student : ", selectedStudents)

	// Slice 사용
	db.Where(ids).Find(&selectedStudents)
	fmt.Println("[6] Slice ConditionedSelected Student : ", selectedStudents)

	deleteSampleData(db)
}

func queryWithSort(db *gorm.DB) {
	fmt.Println("\n\n#### START queryWithSort ####")

	insertSampleData(db)

	// Manty2 인 데이터만 조회
	var selectedStudents []model.Student
	db.Order("name asc, age desc").Where("age > ?", 10).Find(&selectedStudents)
	fmt.Println("[1] ConditionedSortedSelected Student : \n", selectedStudents)

	// Manty2 가 아닌 데이터만 조회
	db.Order("name asc").Order("age desc").Find(&selectedStudents)
	fmt.Println("[2] SortedSelected Student : \n", selectedStudents)

	// sort 할 컬럼을 동적으로 바꿔야 할때.
	db.Clauses(clause.OrderBy{
		Expression: clause.Expr{
			SQL:                "?, ?",
			Vars:               []interface{}{"name asc", "age"},
			WithoutParentheses: true,
		},
	}).Find(&selectedStudents)
	fmt.Println("[3] SortedSelected Student : \n", selectedStudents)

	deleteSampleData(db)
}

func queryWithLimitAndOffset(db *gorm.DB) {
	fmt.Println("\n\n#### START queryWithLimitAndOffset ####")

	insertSampleData(db)

	// 3개 데이터만 조회
	var selectedStudents []model.Student
	db.Order("name asc, age desc").Where("age > ?", 10).Limit(3).Find(&selectedStudents)
	fmt.Println("[1] ConditionedSortedLimitedSelected Student : \n", selectedStudents)

	// Limit 를 취소할때.
	db.Limit(3).Order("name asc").Order("age desc").Limit(-1).Find(&selectedStudents)
	fmt.Println("[2] ConditionedSortedLimitedSelected Student : \n", selectedStudents)

	// Limit 와 offset 을 함께 사용
	db.Limit(3).Offset(5).Order("name asc").Order("age desc").Find(&selectedStudents)
	fmt.Println("[3] ConditionedSortedLimitedSelected Student : \n", selectedStudents)

	deleteSampleData(db)
}
