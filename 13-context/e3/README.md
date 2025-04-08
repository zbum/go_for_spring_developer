## WithTimeout
* 특정시간 뒤에 자동 취소할 수 있는 Context 입니다.
* context.WithTimeout() 함수로 Context 를 생성할 수 있습니다.
* 두번재 반환 값은 취소를 위한 함수 입니다.

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx := context.Background()
    ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
    defer cancel()

    go func() {
        fmt.Println("Starting...")
        select {
        case <-time.After(2 * time.Second):
            fmt.Println("Completed")
        case <-ctx.Done():
            fmt.Println("Cancelled: ", ctx.Err())
        }
    }()

    time.Sleep(3 * time.Second)

}
```


## 참고 링크
* https://youtu.be/mfgBhGu5pco?si=CH22Yl04fLP1q_tP