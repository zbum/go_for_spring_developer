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
	"log"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	wg := sync.WaitGroup{}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			work(ctx, i)
			wg.Done()
		}(i)
	}

	time.AfterFunc(4*time.Second, cancel) // 4초 후에 cancel 을 호출
	wg.Wait()
	log.Println("completed")
}

func work(ctx context.Context, i int) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	slowFn(ctx, i)
}

func slowFn(ctx context.Context, i int) {
	ctx = context.WithValue(ctx, "one", 1)
	ctx = context.WithValue(ctx, "two", 2)

	log.Printf("slow function %d started. \n", i)
	select {
	case <-time.Tick(3 * time.Second):
		log.Printf("slow function %d finished\n", i)
	case <-ctx.Done():
		log.Printf("slow function %d too slow: %s \n", i, ctx.Err())
	}

}

```