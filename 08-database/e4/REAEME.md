# Gorm 으로 Join 데이터 조회
## Join!! 
```go
type Students struct {
    ID     uint `gorm:"primarykey"`
    Name   string
    Age    uint
    Scores []Score `gorm:"foreignKey:StudentID;references:ID"`
}
```

```go
type Scores struct {
    ID        uint
    Score     uint
    StudentID uint
}
```

## Eager Loading
```go
	var selectedStudents []model.Student
	db.Model(&model.Student{}).
		Preload("Scores").
		Where("age = ?", 15).
		Find(&selectedStudents)
```

* Preload 를 했을때 호출되는 SQL
```sql
SELECT * FROM `Students` WHERE age = 15;
SELECT * FROM `Scores` WHERE `Scores`.`student_id` IN (272,274,276,278,280);
```

## Join 조건을 지정
```go
	db.Model(&model.Score{}).
		Select("Students.id, Students.age, Scores.score").
		Joins("left join Students on Scores.student_id = Students.id").
		Where("Students.age = ?", 15).
		Scan(&resultOfStudentWithScores)
```
* Join 을 했을때 호출되는 SQL
```sql
SELECT Students.id, Students.age, Scores.score FROM `Scores` left join Students on Scores.student_id = Students.id WHERE Students.age = 15;
```


## 출처
* gorm : https://gorm.io/ko_KR/docs
