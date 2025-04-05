# Gorm 으로 Join 데이터 조회
## Join!! 
```go
type Students struct {
    ID     uint `gorm:"primarykey"`
    Name   string
    Age    uint
    Scores []Score
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
		Where("name = ?", "Manty1").
		Find(&selectedStudents)
```

## 

db.Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&result{})
// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id

rows, err := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
for rows.Next() {
  ...
}

db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)

// multiple joins with parameter
db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jibum.jung@gmail.com").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "123456789").Find(&user)

```



## 출처
* gorm : https://gorm.io/ko_KR/docs
