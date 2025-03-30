## 슬라이스에서 데이터 읽기 (E6)
* range를 사용하면 slice map 등의 데이터 타입 처리에 반복코드를 제거할 수 있습니다.
```go
  aSlice := []string{"tesla", "nvidia", "apple", "microsoft"}

  for i, v := range aSlice {
    fmt.Printf("Index : %d, Value : %s\n", i, v)
  }
```

## Workshop(e6/w1)
* 제공한 코드의 slice 를 활용하여  2단 부터 9단까지 표시하는 구구단 프로그램을 작성하세요.
* 예상 결과
```
2 x 1 = 2 
2 x 2 = 4 
2 x 3 = 6 
2 x 4 = 8 
2 x 5 = 10 
2 x 6 = 12 
2 x 7 = 14 
2 x 8 = 16 
2 x 9 = 18 
3 x 1 = 3 
// 중간생략....
8 x 9 = 72 
9 x 1 = 9 
9 x 2 = 18 
9 x 3 = 27 
9 x 4 = 36 
9 x 5 = 45 
9 x 6 = 54 
9 x 7 = 63 
9 x 8 = 72 
9 x 9 = 81 
```