package database

import (
	"github.com/zbum/scouter-agent-golang/scouterx/strace"
	"go_for_spring_developer/08-monitoring/e1/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Datasource struct {
	db *gorm.DB
}

func NewDatasource() *Datasource {
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

	// TODO 3 데이터베이스 Query 를 추적하기 위한 플러그인 추가
	db.Use(strace.GormDbPlugin{})

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

	return &Datasource{db}
}

func (d *Datasource) GetDB() *gorm.DB {
	return d.db
}
