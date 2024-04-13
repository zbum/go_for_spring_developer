## Gorm
* Go 언어의 ORM 라이브러리 입니다. 

## 개요
* 완전한 기능을 가진 ORM (지향?)
* Associations (Has One, Has Many, Belongs To, Many To Many, Polymorphism, Single-table inheritance)
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

## CRUD
* 레코드 생성
```go
user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

result := db.Create(&user) // 생성할 데이터의 포인터 넘기기

user.ID             // 입력된 데이터의 primary key를 반환합니다
result.Error        // 에러를 반환합니다
result.RowsAffected // 입력된 레코드의 개수를 반환합니다.
```


