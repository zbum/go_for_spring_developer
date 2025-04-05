package db

import (
	"go_for_spring_developer/08-database/01-common/model"
	"go_for_spring_developer/08-database/01-common/model_with_gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func InitGorm() *gorm.DB {
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
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
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

	// 테이블 자동 생성 + loglevel 조정
	db.Logger.LogMode(logger.Silent)
	err = db.AutoMigrate(
		&model.Student{},
		&model.Score{},
		&model_with_gorm.StudentWithGormModel{},
		&model_with_gorm.ScoreWithGormModel{},
	)
	//db.Logger.LogMode(logger.Info)
	if err != nil {
		panic(err)
	}

	return db
}
