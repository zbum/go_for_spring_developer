package main

import (
	"fmt"
	"go_for_spring_developer/08-database/e2/crud"
	"go_for_spring_developer/08-database/e2/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func main() {
	db := initGorm()
	querySingle(db)
	queryByCondition(db)

}

func initGorm() *gorm.DB {
	cfg := mysql.Config{
		DSN: "root:test@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local",
	}
	var err error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.New(cfg), &gorm.Config{
		Logger: newLogger,
	})
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
	err = db.AutoMigrate(&model.Student{}, &model.Score{})
	if err != nil {
		panic(err)
	}

	return db
}

func querySingle(db *gorm.DB) {
	fmt.Println("\n\n#### START querySingle ####")

	// 생성
	id, insertedCount := crud.Insert(db, &model.Student{Name: "Manty1"})
	fmt.Println("[1] Inserted ID, Count : ", id, insertedCount)

	id, insertedCount = crud.Insert(db, &model.Student{Name: "Manty2"})
	fmt.Println("[2] Inserted ID, Count : ", id, insertedCount)

	// 아무 데이터나 1개 조회 (pk 로 정렬)
	var selectedStudent model.Student
	db.First(&selectedStudent)
	fmt.Println("[3] SingleSelected Student : ", selectedStudent.ID, selectedStudent.Name)

	// 아무 데이터나 1개 조회 (정렬 없음)
	db.Take(&selectedStudent)
	fmt.Println("[3] SingleSelected Student : ", selectedStudent.ID, selectedStudent.Name)

	// 모든 데이터 삭제
	deletedCount := crud.DeleteAll(db)
	fmt.Println("[6] Deleted Count : ", deletedCount)
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
	fmt.Println("[3] ConditionedSelected Student : ", selectedStudent.ID, selectedStudent.Name)

	// Manty2 인 데이터만 조회 (정렬 없음)
	db.Take(&selectedStudent, "Name = ?", "Manty2")
	fmt.Println("[3] ConditionedSelected Student : ", selectedStudent.ID, selectedStudent.Name)

	// 모든 데이터 삭제
	deletedCount := crud.DeleteAll(db)
	fmt.Println("[6] Deleted Count : ", deletedCount)
}
