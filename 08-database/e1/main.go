package main

import (
	"context"
	"fmt"
	"go_for_spring_developer/07-database/e1/crud"
	"go_for_spring_developer/07-database/e1/model"
	"go_for_spring_developer/07-database/e1/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func main() {

	db := initGorm()
	studentRepository := repository.NewStudentRepository()
	studentRepositoryWithContext := repository.NewStudentRepositoryWithContext()
	scoreRepository := repository.NewScoreRepository()

	runCrud(db)
	_ = runRepository(studentRepository, db)
	runRepositoryWithTx(studentRepository, scoreRepository, db)
	runRepositoryWithTxAndContext(studentRepositoryWithContext, db)
}

func initGorm() *gorm.DB {
	cfg := mysql.Config{
		DSN: "root:test@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local",
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

func runCrud(db *gorm.DB) {
	fmt.Println("#### START runCrud ####")
	// 저장 (Insert or Update)
	student := model.Student{Name: "Manty0"}
	saveId, savedCount := crud.Save(db, &student)
	fmt.Println("[0] Inserted ID, Count : ", saveId, savedCount)

	// 생성
	id, insertedCount := crud.Insert(db, &model.Student{Name: "Manty1"})
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

func runRepository(repository *repository.StudentRepository, db *gorm.DB) error {
	fmt.Println("#### START runRepository ####")
	// 저장 (Insert or Update)
	student := model.Student{Name: "Manty0"}
	if saveId, savedCount, err := repository.Save(db, &student); err != nil {
		return err
	} else {
		fmt.Println("[0] Inserted ID, Count : ", saveId, savedCount)
	}

	// 생성
	var id uint
	if id, insertedCount, err := repository.Insert(db, &model.Student{Name: "Manty1"}); err != nil {
		return err
	} else {
		fmt.Println("[1] Inserted ID, Count : ", id, insertedCount)
	}

	if selectedStudent, err := repository.FindById(db, id); err != nil {
		return err
	} else {
		fmt.Println("[2] Selected Student : ", selectedStudent)
	}

	// 단건 수정
	if _, err := repository.UpdateNameById(db, id, "Manty2"); err != nil {
		return err
	}

	if selectedStudent, err := repository.FindById(db, id); err != nil {
		return err
	} else {
		fmt.Println("[3] Selected Student : ", selectedStudent)
	}

	// 전체 데이터 수정
	if _, err := repository.UpdateNames(db, "Manty3"); err != nil {
		return err
	}

	if allStudents, err := repository.FindAll(db); err != nil {
		return err
	} else {
		fmt.Println("[4] allStudents : ", allStudents)
	}

	// id 데이터 삭제
	if deletedCount, err := repository.DeleteById(db, id); err != nil {
		return err
	} else {
		fmt.Println("[5] Deleted ID, Count : ", id, deletedCount)
	}

	// 모든 데이터 삭제
	if deletedCount, err := repository.DeleteAll(db); err != nil {
		return err
	} else {
		fmt.Println("[6] Deleted Count : ", deletedCount)
	}
	return nil
}

func runRepositoryWithTx(repository *repository.StudentRepository, scoreRepository *repository.ScoreRepository, db *gorm.DB) {
	fmt.Println("#### START runRepositoryWithTx ####")

	// 기본 트랜젝션을 일시적으로 중단합니다.
	db.Session(&gorm.Session{SkipDefaultTransaction: true})
	db.Transaction(func(tx *gorm.DB) error {
		// 저장 (Insert or Update)
		student := model.Student{Name: "Manty0"}
		if saveId, savedCount, err := repository.Save(tx, &student); err != nil {
			return err
		} else {
			fmt.Println("[0] Inserted ID, Count : ", saveId, savedCount)
		}

		score := model.Score{Score: 99, StudentID: student.ID}
		if saveId, savedCount, err := scoreRepository.Save(tx, &score); err != nil {
			return err
		} else {
			fmt.Println("[0] Inserted ID, Count : ", saveId, savedCount)
		}

		// 생성
		var studentId uint
		if id, insertedCount, err := repository.Insert(tx, &model.Student{Name: "Manty1"}); err != nil {
			return err
		} else {
			studentId = id
			fmt.Println("[1] Inserted ID, Count : ", id, insertedCount)
		}

		if selectedStudent, err := repository.FindById(tx, studentId); err != nil {
			return err
		} else {
			fmt.Println("[2] Selected Student : ", selectedStudent)
		}

		// 단건 수정
		if _, err := repository.UpdateNameById(tx, studentId, "Manty2"); err != nil {
			return err
		}

		if selectedStudent, err := repository.FindById(tx, studentId); err != nil {
			return err
		} else {
			fmt.Println("[3] Selected Student : ", selectedStudent)
		}

		// 전체 데이터 수정
		if _, err := repository.UpdateNames(tx, "Manty3"); err != nil {
			return err
		}

		if allStudents, err := repository.FindAll(tx); err != nil {
			return err
		} else {
			fmt.Println("[4] allStudents : ", allStudents)
		}

		// id 데이터 삭제
		if deletedCount, err := repository.DeleteById(tx, studentId); err != nil {
			return err
		} else {
			fmt.Println("[5] Deleted ID, Count : ", studentId, deletedCount)
		}

		// 모든 데이터 삭제
		if deletedCount, err := repository.DeleteAll(tx); err != nil {
			return err
		} else {
			fmt.Println("[6] Deleted Count : ", deletedCount)
		}

		return nil
	})
}

func runRepositoryWithTxAndContext(repository *repository.StudentRepositoryWithContext, db *gorm.DB) {
	fmt.Println("#### START runRepositoryWithTxAndContext ####")

	background := context.Background()
	ctx, _ := context.WithTimeout(background, 10000*time.Millisecond)

	// 기본 트랜젝션을 일시적으로 중단합니다.
	db.Session(&gorm.Session{SkipDefaultTransaction: true})
	db.Transaction(func(tx *gorm.DB) error {
		// 저장 (Insert or Update)
		student := model.Student{Name: "Manty0"}
		if saveId, savedCount, err := repository.Save(ctx, tx, &student); err != nil {
			return err
		} else {
			fmt.Println("[0] Inserted ID, Count : ", saveId, savedCount)
		}

		// 생성
		var studentId uint
		if id, insertedCount, err := repository.Insert(ctx, tx, &model.Student{Name: "Manty1"}); err != nil {
			return err
		} else {
			studentId = id
			fmt.Println("[1] Inserted ID, Count : ", id, insertedCount)
		}

		if selectedStudent, err := repository.FindById(ctx, tx, studentId); err != nil {
			return err
		} else {
			fmt.Println("[2] Selected Student : ", selectedStudent)
		}

		// 단건 수정
		if _, err := repository.UpdateNameById(ctx, tx, studentId, "Manty2"); err != nil {
			return err
		}

		if selectedStudent, err := repository.FindById(ctx, tx, studentId); err != nil {
			return err
		} else {
			fmt.Println("[3] Selected Student : ", selectedStudent)
		}

		// 전체 데이터 수정
		if _, err := repository.UpdateNames(ctx, tx, "Manty3"); err != nil {
			return err
		}

		if allStudents, err := repository.FindAll(ctx, tx); err != nil {
			return err
		} else {
			fmt.Println("[4] allStudents : ", allStudents)
		}

		// id 데이터 삭제
		if deletedCount, err := repository.DeleteById(ctx, tx, studentId); err != nil {
			return err
		} else {
			fmt.Println("[5] Deleted ID, Count : ", studentId, deletedCount)
		}

		// 모든 데이터 삭제
		if deletedCount, err := repository.DeleteAll(ctx, tx); err != nil {
			return err
		} else {
			fmt.Println("[6] Deleted Count : ", deletedCount)
		}

		return nil
	})
}
