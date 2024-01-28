package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "test",
		Net:                  "tcp",
		Addr:                 "localhost:13306",
		DBName:               "test",
		AllowNativePasswords: true,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected")

	students, err := findStudentsByName("Zbum");
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", students)
}

type Student struct {
	ID       int64
	Name     string
	Nickname string
	Score    float32
}

func findStudentsByName(name string) ([]Student, error) {
	// An students slice to hold data from returned rows.
	var students []Student

	rows, err := db.Query("SELECT * FROM Students WHERE name = ?", name)
	if err != nil {
		return nil, fmt.Errorf("findStudentsByName %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.ID, &student.Name, &student.Nickname, &student.Score); err != nil {
			return nil, fmt.Errorf("findStudentsByName %q: %v", name, err)
		}
		students = append(students, student)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("findStudentsByName %q: %v", name, err)
	}
	return students, nil
}
