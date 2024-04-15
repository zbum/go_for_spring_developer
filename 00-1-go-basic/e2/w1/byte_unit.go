package main

import "fmt"

// TODO MiB, GiB, TiB 상수를 선언해 주세요.
const (
	_           = iota
	KiB float64 = 1 << (iota * 10)
)

func main() {
	var a float64 = 10000000000
	fmt.Println(humanReadable(a))

}

// TODO MiB, GiB, TiB 조건도 추가합니다.
func humanReadable(a float64) string {
	if a >= KiB {
		return fmt.Sprintf("%.2fKiB", a/KiB)
	}
	return fmt.Sprintf("%.2fB", a)
}
