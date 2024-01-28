## 스탠다드 라이브러리 사용.
* Go 언어는 데이터베이스 접근을 위한 "database/sql" 라이브러리를 제공합니다. 
* 다만, 데이터베이스 드라이브는 개발사의 드라이브를 사용합니다. 

## 데이터베이스 접속
* 데이터베이스 접속은 `Open(driverName, dataSourceName string) (*DB, error)` 함수를 사용합니다. 
* driverName 은 "github.com/go-sql-driver/mysql" 드라이버가 초기화 할때, 다음의 코드로 등록을 해 두었습니다. 그래서 "mysql" 을 사용하면 됩니다.
```go
func init() {
	sql.Register("mysql", &MySQLDriver{})
}
```
* DSN(dataSourceName)은 다음의 형태로 작성해야 합니다.
```
[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
```
```
username:password@protocol(address)/dbname?param=value
```
* 대부분 드라이버 모듈이 DSN 작성을 지원합니다. MySql 은 다음과 같이 DSN 을 생성할 수 있습니다.
```go
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "test",
		Net:                  "tcp",
		Addr:                 "localhost:13306",
		DBName:               "test",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
```

## SELECT
* 데이터를 조회하기 전에 테이블과 데이터를 먼저 준비합니다. 
```sql
DROP TABLE IF EXISTS Students;
CREATE TABLE Students (
  id         INT AUTO_INCREMENT NOT NULL,
  name      VARCHAR(128) NOT NULL,
  nickname     VARCHAR(255) NOT NULL,
  score      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO Students
  (name, nickname, score)
VALUES
  ('Zbum', 'Manty', 99.55),
  ('DongMin', 'Dongmyo', 63.99);
```

* Student 스트럭트를 다음과 같이 정의 합니다. 
```go
type Student struct {
	ID       int64
	Name     string
	Nickname string
	Score    float32
}
```
* 다음의 코드로 데이터를 조회합니다.
```go
// findStudentsByName queries for students that have the specified student name.
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
```
* Student 슬라이스를 선언했습니다. 데이터베이스에서 조회한 결과 값을 담아 반환하기 위함입니다. 각 스트럭트의 필드는 데이터베이스의 컬럼 이름과 타입에 대응되어야 합니다. 
* DB.Query 함수를 사용하여 SELECT 문을 실행하였습니다. 쿼리의 첫 번째 매개변수는 SQL 문입니다.  매개변수 값에서 SQL 문을 분리하면(예를 들어 fmt.Sprintf와 연결하지 않고) 데이터베이스가 SQL 텍스트와 별도로 값을 보낼 수 있도록 하여 SQL 삽입 위험을 제거할 수 있습니다.여

* 이 함수가 종료할때 rows 를 close 할 수 있도록 defer rows.Close() 를 사용하였습니다.

* Rows.Scan을 사용하여 각 행의 열 값을 Student 구조체 필드에 할당합니다. 

* Scan은 컬럼 값이 기록될 Go 값에 대한 포인터 목록을 가져옵니다. 여기에서는 & 연산자를 사용하여 생성된 student 변수의 필드에 포인터를 전달합니다. 스캔은 포인터를 통해 쓰기를 수행하여 구조체 필드를 업데이트합니다.

* 루프 내에서 열 값을 구조체 필드로 스캔할 때 오류가 있는지 확인하세요.

* 루프 내에서 새 앨범을 앨범 슬라이스에 추가합니다.

* 루프 후에 행.Err을 사용하여 전체 쿼리에서 오류를 확인합니다. 쿼리 자체가 실패하는 경우 여기에서 오류를 확인하는 것이 결과가 불완전하다는 것을 알아내는 유일한 방법입니다.

### 실행하기
* mysql 드라이버 의존성을 가져오기 위해 다음을 실행합니다.
```shell
$ go get .
go get: added github.com/go-sql-driver/mysql v1.6.0
```

* main.go 실행
```shell
$ go run .
```


## 참고 링크
* https://go.dev/doc/tutorial/database-access