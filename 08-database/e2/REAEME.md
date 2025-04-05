## Gorm 으로 데이터 조회
### SELECT (querySingle)
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

tx := db.First(&user)
tx.RowsAffected // 찾아낸 레코드의 개수
tx.Error        // error이나 nil

// ErrRecordNotFound 를 체크하는 방법
errors.Is(result.Error, gorm.ErrRecordNotFound)
```
### SELECT By Primary Key (queryByPrimaryKey)
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
