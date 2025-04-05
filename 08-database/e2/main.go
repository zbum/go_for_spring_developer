package main

import (
	"fmt"
	"go_for_spring_developer/08-database/01-common/crud"
	"go_for_spring_developer/08-database/01-common/db"
	"go_for_spring_developer/08-database/01-common/model"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strconv"
)

func main() {
	db := db.InitGorm()
	querySingle(db)
	queryByPrimaryKey(db)
	queryByCondition(db)
	queryMultipleByCondition(db)
}

func querySingle(db *gorm.DB) {

	fmt.Println("\n\n#### START querySingle ####")
	// loglevel 을 잠시 낮춥니다.
	db.Logger.LogMode(logger.Silent)

	// 테스트 데이터 입력
	id, insertedCount := crud.Insert(db, &model.Student{Name: "Manty1"})
	fmt.Println("[1] Inserted ID, Count : ", id, insertedCount)

	id, insertedCount = crud.Insert(db, &model.Student{Name: "Manty2"})
	fmt.Println("[2] Inserted ID, Count : ", id, insertedCount)

	// pk 로 정렬한 첫번째 데이터 조회
	var selectedStudent model.Student
	tx := db.First(&selectedStudent)
	fmt.Printf("[3] SingleSelected Student : ROWs: %d, ERROR: %v\n", tx.RowsAffected, tx.Error)
	fmt.Println("[3] SingleSelected Student : ", selectedStudent)

	// 아무 데이터나 1개 조회 (정렬 없음)
	db.Take(&selectedStudent)
	fmt.Println("[3] SingleSelected Student : ", selectedStudent)

	// 모든 테스트 데이터 삭제
	deletedCount := crud.DeleteAll(db)
	fmt.Println("[6] Deleted Count : ", deletedCount)

	// loglevel 복구
	db.Logger.LogMode(logger.Info)
}

func queryByPrimaryKey(db *gorm.DB) {
	fmt.Println("\n\n#### START queryByPrimaryKey ####")

	var ids []uint

	// 생성
	id, _ := crud.Insert(db, &model.Student{Name: "Manty1"})
	ids = append(ids, id)

	id, _ = crud.Insert(db, &model.Student{Name: "Manty2"})
	ids = append(ids, id)
	fmt.Println("[1] Inserted IDs, Count : ", ids, len(ids))

	// 마지막 insert한 데이터만 조회
	var selectedStudent model.Student
	db.First(&selectedStudent, id)
	fmt.Println("[2] PrimaryKey Selected Student : ", selectedStudent)

	// 마지막 insert한 데이터만 조회
	// 문자열로 조회해도 동작합니다.
	var selectedStudentByString model.Student
	db.Take(&selectedStudentByString, strconv.FormatUint(uint64(id), 10))
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

	// 모든 데이터 삭제
	deletedCount := crud.DeleteAll(db)
	fmt.Println("[5] Deleted Count : ", deletedCount)
}

func queryByCondition(db *gorm.DB) {
	fmt.Println("\n\n#### START queryByCondition ####")

	// 생성
	id, insertedCount := crud.Insert(db, &model.Student{Name: "Manty1"})
	fmt.Println("[1] Inserted ID, Count : ", id, insertedCount)

	id, insertedCount = crud.Insert(db, &model.Student{Name: "Manty2"})
	fmt.Println("[2] Inserted ID, Count : ", id, insertedCount)

	// Manty2 인 데이터만 조회 (pk 로 정렬)
	var selectedStudent model.Student
	db.First(&selectedStudent, "Name = ?", "Manty2")
	fmt.Println("[3] ConditionedSelected Student : ", selectedStudent)

	// Manty2 인 데이터만 조회 (정렬 없음)
	db.Take(&selectedStudent, "Name = ?", "Manty2")
	fmt.Println("[3] ConditionedSelected Student : ", selectedStudent)

	// 모든 데이터 삭제
	deletedCount := crud.DeleteAll(db)
	fmt.Println("[6] Deleted Count : ", deletedCount)
}

func queryMultipleByCondition(db *gorm.DB) {
	fmt.Println("\n\n#### START queryMultipleByCondition ####")

	// 생성 (Manty1, Manty2)
	id, insertedCount := crud.Insert(db, &model.Student{Name: "Manty1"})
	fmt.Println("[1] Inserted ID, Count : ", id, insertedCount)

	id, insertedCount = crud.Insert(db, &model.Student{Name: "Manty2"})
	fmt.Println("[2] Inserted ID, Count : ", id, insertedCount)

	// 전체 데이터 조회
	var selectedStudents []model.Student
	db.Find(&selectedStudents)
	fmt.Printf("[3] SelectAll Student : \n%v", selectedStudents)

	// 전체 데이터 조회(PK로 정렬)
	db.Order("id").Find(&selectedStudents)
	fmt.Printf("[4] SelectAll Student : \n%v", selectedStudents)

	// 전체 데이터 조회(PK로 역정렬(내림차순))
	db.Order("id desc").Find(&selectedStudents)
	fmt.Printf("[5] SelectAll Student : \n%v", selectedStudents)

	// 모든 데이터 삭제
	deletedCount := crud.DeleteAll(db)
	fmt.Println("[6] Deleted Count : ", deletedCount)
}
