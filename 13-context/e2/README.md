## WithCancel

* 실행을 취소할 수 있는 Context 입니다. 
* context.WithCancel() 함수로 CancelContext 를 생성할 수 있습니다. 
* 두번재 반환 값은 취소를 위한 함수 입니다.
* 부모가 취소하면 자식 Context 도 취소가 됩니다. 
* 취소가 발생하면 Context.Done() 으로 얻을 수 있는 채널에서 값이 발생합니다. 

## 예제
```go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)

    go func() {
        fmt.Println("Starting...")
        select {
        case <-time.After(2 * time.Second):
            fmt.Println("Completed")
        case <-ctx.Done():
            fmt.Println("Cancelled: ", ctx.Err())
        }
    }()

    time.Sleep(1 * time.Second)
    cancel()

    time.Sleep(1 * time.Second)
}
```