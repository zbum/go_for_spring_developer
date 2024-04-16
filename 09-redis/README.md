## go-redis
* go-redis는 redis client로서 다양한 기능을  type-safe API 를 제공합니다.
* https://github.com/redis/go-redis

## 설치
* go-redis/v9 을 설치합니다.
```shell
$ go get github.com/redis/go-redis/v9
```
## 레디스 접속
* 레디스 서버 접속 
```go
import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
    client := redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "", // no password set
        DB:		  0,  // use default DB
    })
}
```
* 다른 방법으로 접속할 수 있습니다.
```go
opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
if err != nil {
	panic(err)
}

client := redis.NewClient(opt)
```

* set , get 을 이용하여 간단한 문자열 입력, 출력 확인
```go
ctx := context.Background()

err := client.Set(ctx, "foo", "bar", 0).Err()
if err != nil {
    panic(err)
}

val, err := client.Get(ctx, "foo").Result()
if err != nil {
    panic(err)
}
fmt.Println("foo", val)
```
* HSet, HGetAll 을 이용하여 map 정보 등록과 조회를 수행합니다.
```go
session := map[string]string{"name": "John", "surname": "Smith", "company": "Redis", "age": "29"}
for k, v := range session {
    err := client.HSet(ctx, "user-session:123", k, v).Err()
    if err != nil {
        panic(err)
    }
}

userSession := client.HGetAll(ctx, "user-session:123").Val()
fmt.Println(userSession)
```

## Redis Cluster 접속
* Redis Cluster 에 접속하기 위해서는 redis.NewClusterClient 생성자 함수를 사용해야 합니다. 
```go
client := redis.NewClusterClient(&redis.ClusterOptions{
    Addrs: []string{":16379", ":16380", ":16381", ":16382", ":16383", ":16384"},

    // To route commands by latency or randomly, enable one of the following.
    //RouteByLatency: true,
    //RouteRandomly: true,
})
```

##  workshop (w1) 
* workshop.go 를 활용하여 ZSET, ZRANGE 로 sorted set 데이터를 입력, 조회하는 코드를 작성하세요.

## 출처
* https://redis.io/docs/latest/develop/connect/clients/go/