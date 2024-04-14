#### 슬라이스에서 데이터 읽기 (E6)
```go
  aSlice := []string{"tesla", "nvidia", "apple", "microsoft"}

  for i, v := range aSlice {
  fmt.Printf("Index : %d, Value : %s\n", i, v)
  }
```