## Gorm 으로 데이터 모델

* gorm에서 제공하는 gorm.DeletedAt를 사용한다면 조회 쿼리가 바뀝니다.
```go
type User struct {
  ID           string `gorm:"primarykey;size:16"`
  Name         string `gorm:"size:24"`
  DeletedAt    gorm.DeletedAt `gorm:"index"`
}

var user = User{ID: 15}
db.First(&user)
//  SELECT * FROM `users` WHERE `users`.`id` = '15' AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
```

* gorm에서 제공하는 gorm.Model을 임베딩 하면 id, created_at, updated_at, deleted_at 컬럼이 자동으로 관리 됩니다. 
```go
type StudentWithGormModel struct {
    gorm.Model
    Name   string
    Scores []ScoreWithGormModel `gorm:"foreignKey:StudentID;references:ID"`
}
```

```
mysql> desc StudentsWithGormModel;
+------------+---------------------+------+-----+---------+----------------+
| Field      | Type                | Null | Key | Default | Extra          |
+------------+---------------------+------+-----+---------+----------------+
| id         | bigint(20) unsigned | NO   | PRI | NULL    | auto_increment |
| created_at | datetime(3)         | YES  |     | NULL    |                |
| updated_at | datetime(3)         | YES  |     | NULL    |                |
| deleted_at | datetime(3)         | YES  | MUL | NULL    |                |
| name       | longtext            | YES  |     | NULL    |                |
+------------+---------------------+------+-----+---------+----------------+
5 rows in set (0.00 sec)

```


## 출처
* gorm : https://gorm.io/ko_KR/docs
