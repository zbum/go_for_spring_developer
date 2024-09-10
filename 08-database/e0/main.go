package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Student struct {
	ID   int64
	Name string
}

func main() {

	db := initDatasource()

	err := insertStudent(db, &Student{ID: 1, Name: "Manty"})
	if err != nil {
		log.Fatal(err)
	}
	students, err := findStudentsByName(db, "Manty")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", students)
}

func initDatasource() *sql.DB {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "test",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "test",
		AllowNativePasswords: true,
	}
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected")

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(1 * time.Hour) // idle 상태로 유지되는 시간
	db.SetConnMaxLifetime(1 * time.Hour) // connection의 재사용 가능 시간

	return db
}

func insertStudent(db *sql.DB, student *Student) error {

	_, err := db.Query("INSERT INTO Students(id, name) VALUES (?,?)", student.ID, student.Name)
	if err != nil {
		return fmt.Errorf("insertStudent %v: %v", student, err)
	}
	return nil
}

func findStudentsByName(db *sql.DB, name string) ([]Student, error) {
	// An students slice to hold data from returned rows.
	var students []Student

	rows, err := db.Query("SELECT id, name FROM Students WHERE name = ?", name)
	if err != nil {
		return nil, fmt.Errorf("findStudentsByName %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.ID, &student.Name); err != nil {
			return nil, fmt.Errorf("findStudentsByName %q: %v", name, err)
		}
		students = append(students, student)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("findStudentsByName %q: %v", name, err)
	}
	return students, nil
}
