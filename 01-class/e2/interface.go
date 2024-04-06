package main

import (
	"fmt"
	"sort"
)

type SortableSlice []int

func (s SortableSlice) Len() int {
	return len(s)
}

func (s SortableSlice) Less(i, j int) bool {
	return i > j
}

func (s SortableSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func main() {
	data1 := SortableSlice([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	sort.Sort(data1)

	fmt.Println(data1)
}
