

### 반복문
* Go는 반복문으로 for 키워드만 제공합니다. (while 이 없습니다.)
#### 구구단의 2단 (E5)
```go
package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Printf("%d X %d = %d \n", 2, i, 2*i)
	}
}
```
* range를 사용하면 slice map 등의 데이터 타입 처리에 반복코드를 제거할 수 있습니다.
