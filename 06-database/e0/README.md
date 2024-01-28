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
   
