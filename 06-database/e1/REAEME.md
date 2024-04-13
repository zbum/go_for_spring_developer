## Gorm
* Go 언어의 대표적인 ORM 라이브러리 입니다.
* Go 언어 1.0의 출시(2012) 출시 후 2년 뒤에 시작한 라이브러리 입니다.
![img.png](img.png)

## 개요
* 완전한 기능을 가진 ORM (지향?)
* 연관관계 지원 
  * Has One(one-to-one)
  * Has Many(one-to-many)
  * Belongs To(manty-to-one)
  * Many To Many
  * Polymorphism
  * Single-table inheritance
* Hooks (Before/After Create/Save/Update/Delete/Find)
* Preload, Joins를 통한 데이터 가져오기
* Transactions, Nested Transactions, Save Point, RollbackTo to Saved Point
* Context, Prepared Statement 모드, DryRun 모드
* Batch Insert, FindInBatches, Find/Create with Map, CRUD with SQL Expr and Context Valuer
* SQL Builder, Upsert, Locking, Optimizer/Index/Comment Hints, Named Argument, SubQuery
* Composite Primary Key, Indexes, Constraints
* Auto Migrations
* Logger
* 확장 가능하고 유연한 플러그인 API: Database Resolver (다중 데이터베이스, 읽기 / 쓰기 분할) / Prometheus

## 설치
```go
go get -u gorm.io/gorm
go get -u github.com/go-sql-driver/mysql
```

## 준비
```sql
create database gorm;
```

## 설정
* GORM은 MySQL, PostgreSQL, SQLite, SQL Server, TiDB 를 지원합니다.  
* GORM이 사용할 데이터소스를 설정은 전체 시스템에서 1회만 실행해야 합니다. 
```go
package main

import (
    "context"
    "fmt"
    "gorm.io/driver/mysql" //"github.com/go-sql-driver/mysql" 를 import 하지 않습니다. 내부에서 mysql 드라이버를 사용하고 있습니다.
    "gorm.io/gorm"
    "time"
)

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
	err = db.AutoMigrate(&model.Student{})
	if err != nil {
		panic(err)
	}

	return db
}
```

## CRUD
### INSERT
* 레코드 생성은 Create 함수를 사용합니다. 
* 매개변수의 타입은 value interface{} 이므로 아무 타입이나 입력할 수 있지만 처리 후에 데이터를 변경해야 하므로 반드시 포인터를 전달해야 합니다. 
* crud.go Insert 함수
```go
func Insert(db *gorm.DB, student *model.Student) (id uint, rowsAffected int64) {
    tx := db.Create(student) 
    if tx.Error != nil { // 에러를 반환합니다
        fmt.Println(db.Error)
    }
    return student.ID, tx.RowsAffected // 입력된 데이터의 primary key, 레코드 개수를 반환합니다
}
```
* 만약, 여러개의 레코드를 Insert 하고자 한다면 포인터 슬라이스를 매개 변수로 넘겨 주면 됩니다.
```go
func Inserts(db *gorm.DB, students []*model.Student) (rowsAffected int64) {
	tx := db.Create(students)
	if tx.Error != nil {
		fmt.Println(db.Error)
	}
	return tx.RowsAffected
}
```

### UPDATE - SAVE
* Save는 변수 필드에 프라이머리 키 값이 포함되어 있으면 업데이트를 처리하고 만약 빈값이면 인서트를 처리합니다. 
```go
db.Save(&User{Name: "zbum", Age: 100})
// INSERT INTO `users` (`name`,`age`,`birthday`,`update_at`) VALUES ("zbum",100,"0000-00-00 00:00:00","0000-00-00 00:00:00")

db.Save(&User{ID: 1, Name: "zbum", Age: 100})
// UPDATE `users` SET `name`="zbum",`age`=100,`birthday`="0000-00-00 00:00:00",`update_at`="0000-00-00 00:00:00" WHERE `id` = 1
```

### UPDATE
* 단일 레코드를 수정할때는 Update 메소드를 사용합니다. Update 는 Model 의 값에 기본 키값이 있는 경우, 그 키를 사용합니다. 
* 만약 Update 시에 조건절이 없다면 ErrMissingWhereClause 가 발생합니다.
```go
// Update with conditions
db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;

// User's ID is `111`:
db.Model(&user).Update("name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

// Update with conditions and model value
db.Model(&user).Where("active = ?", true).Update("name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;
```

* 구조체 또는 맵을 이용하여 업데이트 가능합니다. 구조체를 사용한 경우, 제로벨류는 사용하지 않습니다. 
```go
// Update attributes with `struct`, will only update non-zero fields
db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

// Update attributes with `map`
db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

```

### DELETE
* 삭제는 반드시 primary 키가 값이 있어야 동작합니다. 만약 primary 키 없이 여러건을 삭제할 때는 BATCH DELETE를 사용합니다.
```go
db.Delete(&User{}, 10)
// DELETE FROM users WHERE id = 10;

db.Delete(&User{}, "10")
// DELETE FROM users WHERE id = 10;

db.Delete(&users, []int{1,2,3})
// DELETE FROM users WHERE id IN (1,2,3);

```

### BATCH DELETE
```go
db.Where("email LIKE ?", "%jinzhu%").Delete(&Email{})
// DELETE from emails where email LIKE "%jinzhu%";

db.Delete(&Email{}, "email LIKE ?", "%jinzhu%")
// DELETE from emails where email LIKE "%jinzhu%";
```

### Soft DELETE
* 필드에 gorm.DeletedAt 이 포함되어 있다면 자동으로 Soft DELETE 가 처리됩니다.
* Delete 메소드가 호출되면 실제 데이터를 삭제하지 않고 DeletedAt 필드에 현재시작을 등록합니다.
* 일반적인 Query 메소드로는 해당 레코드를 조회할 수 없습니다.

* Soft DELETE 한 레코드 조회하기
```go
db.Unscoped().Where("age = 20").Find(&users)
// SELECT * FROM users WHERE age = 20;
```

* Soft DELETE 레코드 삭제하기
```go
db.Unscoped().Delete(&order)
// DELETE FROM orders WHERE id=10;
```

### SELECT
* 단건 조회를 위해 First, Take, Last 메소드를 사용할 수 있습니다. 
* 이 메서드는 쿼리문에 LIMIT 1 구문을 추가하고, 만일 해당하는 레코드가 없을 경우 ErrRecordNotFound 에러를 반환합니다.
```go
// primary key로 정렬하여 첫번째 레코드를 가져 옵니다.
db.First(&user)
// SELECT * FROM users ORDER BY id LIMIT 1;

// 정렬 순서와 상관 없이 레코드 하나를 가져 옵니다.
db.Take(&user)
// SELECT * FROM users LIMIT 1;

// primary key로 내림차순 정렬한 첫번째 값을 가져 옵니다.
db.Last(&user)
// SELECT * FROM users ORDER BY id DESC LIMIT 1;

result := db.First(&user)
result.RowsAffected // 찾아낸 레코드의 개수
result.Error        // error이나 nil

// ErrRecordNotFound 를 체크하는 방법
errors.Is(result.Error, gorm.ErrRecordNotFound)
```
### SELECT By Primary Key
* primary Key가 숫자인 경우, 인라인 조건을 사용하면 기본 키를 사용하여 개체를 검색할 수 있습니다. 
* 문자열 작업 시 SQL 주입을 방지하려면 각별한 주의가 필요합니다. 
```go
db.First(&user, 10)
// SELECT * FROM users WHERE id = 10;

db.First(&user, "10")
// SELECT * FROM users WHERE id = 10;

db.Find(&users, []int{1,2,3})
// SELECT * FROM users WHERE id IN (1,2,3);
```

* primary Key가 문자열이면, 쿼리는 다음과 같이 생성됩니다.
```go
db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";
```

* 만약 목적지 변수에 primary key 값이 할당되어 있다면 그 값을 사용합니다.
```go
var user = User{ID: 10}
db.First(&user)
// SELECT * FROM users WHERE id = 10;

var result User
db.Model(User{ID: 10}).First(&result)
// SELECT * FROM users WHERE id = 10;
```
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

### SELECT ALL
```go
// Get all records
result := db.Find(&users)
// SELECT * FROM users;

result.RowsAffected // returns found records count, equals `len(users)`
result.Error        // returns error
```

### SELECT 조건
```go
// Get first matched record
db.Where("name = ?", "zbum").First(&user)
// SELECT * FROM users WHERE name = 'zbum' ORDER BY id LIMIT 1;

// Get all matched records
db.Where("name <> ?", "zbum").Find(&users)
// SELECT * FROM users WHERE name <> 'zbum';

// IN
db.Where("name IN ?", []string{"zbum", "zbum 2"}).Find(&users)
// SELECT * FROM users WHERE name IN ('zbum','zbum 2');

// LIKE
db.Where("name LIKE ?", "%jin%").Find(&users)
// SELECT * FROM users WHERE name LIKE '%jin%';

// AND
db.Where("name = ? AND age >= ?", "zbum", "22").Find(&users)
// SELECT * FROM users WHERE name = 'zbum' AND age >= 22;

// Time
db.Where("updated_at > ?", lastWeek).Find(&users)
// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

// BETWEEN
db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

```

### 구조체, map SELECT 조건
```go
// 구조체
db.Where(&User{Name: "zbum", Age: 20}).First(&user)
// SELECT * FROM users WHERE name = "zbum" AND age = 20 ORDER BY id LIMIT 1;

// Map
db.Where(map[string]interface{}{"name": "zbum", "age": 20}).Find(&users)
// SELECT * FROM users WHERE name = "zbum" AND age = 20;

// Slice of primary keys
db.Where([]int64{20, 21, 22}).Find(&users)
// SELECT * FROM users WHERE id IN (20, 21, 22);
```

> 구조체에서 제로값 (0, '', false)이 있으면 조건에서 제외됩니다. 만약 제로값을 조건에 포함하고 싶다면 map을 사용해야 합니다.
 
```go
db.Where(&User{Name: "zbum", Age: 0}).Find(&users)
// SELECT * FROM users WHERE name = "zbum";

db.Where(map[string]interface{}{"Name": "zbum", "Age": 0}).Find(&users)
// SELECT * FROM users WHERE name = "zbum" AND age = 0;
```

### 정렬
```go
db.Order("age desc, name").Find(&users)
// SELECT * FROM users ORDER BY age desc, name;

// Multiple orders
db.Order("age desc").Order("name").Find(&users)
// SELECT * FROM users ORDER BY age desc, name;

db.Clauses(clause.OrderBy{
  Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
}).Find(&User{})
// SELECT * FROM users ORDER BY FIELD(id,1,2,3)
```

### Limit, Offset
```go
db.Limit(3).Find(&users)
// SELECT * FROM users LIMIT 3;

// Cancel limit condition with -1
db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
// SELECT * FROM users LIMIT 10; (users1)
// SELECT * FROM users; (users2)

db.Offset(3).Find(&users)
// SELECT * FROM users OFFSET 3;

db.Limit(10).Offset(5).Find(&users)
// SELECT * FROM users OFFSET 5 LIMIT 10;

// Cancel offset condition with -1
db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
// SELECT * FROM users OFFSET 10; (users1)
// SELECT * FROM users; (users2)
```

### Join!!
```go
type result struct {
  Name  string
  Email string
}

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



