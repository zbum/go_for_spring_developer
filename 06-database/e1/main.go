package main

import (
	"fmt"
	"go_for_spring_developer/06-database/e1/crud"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {

	db := initGorm()

	// 저장 (Insert or Update)
	student := crud.Student{Name: "Manty0"}
	saveId, savedCount := crud.Save(db, &student)
	fmt.Println("[0] Inserted ID, Count : ", saveId, savedCount)

	// 생성
	id, insertedCount := crud.Insert(db, &crud.Student{Name: "Manty1"})
	fmt.Println("[1] Inserted ID, Count : ", id, insertedCount)

	selectedStudent := crud.FindById(db, id)
	fmt.Println("[2] Selected Student : ", selectedStudent)

	// 단건 수정
	crud.UpdateNameById(db, id, "Manty2")

	selectedStudent = crud.FindById(db, id)
	fmt.Println("[3] Selected Student : ", selectedStudent)

	// 전체 데이터 수정
	crud.UpdateNames(db, "Manty3")

	allStudents := crud.FindAll(db)
	fmt.Println("[4] allStudents : ", allStudents)

	// id 데이터 삭제
	deletedCount := crud.DeleteById(db, id)
	fmt.Println("[5] Deleted ID, Count : ", id, deletedCount)

	// 모든 데이터 삭제
	deletedCount = crud.DeleteAll(db)
	fmt.Println("[6] Deleted Count : ", deletedCount)
}

func initGorm() *gorm.DB {
	cfg := mysql.Config{
		DSN: "root:test@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local",
	}
	var err error

	db, err := gorm.Open(mysql.New(cfg), &gorm.Config{})
	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDb.SetMaxIdleConns(100)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxIdleTime(1 * time.Hour) // idle 상태로 유지되는 시간
	sqlDb.SetConnMaxLifetime(1 * time.Hour) // connection의 재사용 가능 시간

	// 테이블 자동 생성
	err = db.AutoMigrate(&crud.Student{})
	if err != nil {
		panic(err)
	}

	return db
}
