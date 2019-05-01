package main

import (
	"fmt"
)

func main() {
	result := SumTwoNum([]int{1, 3, 2, 5, 7, 11, 15}, 9)
	fmt.Println(result)
}

func SumTwoNum(arrayInt []int, target int) []int {
	index := make(map[int]int, len(arrayInt))
	for i, d := range arrayInt {
		if _, ok := index[target-arrayInt[i]]; ok {
			return []int{index[target-arrayInt[i]], i}
		}
		index[d] = i
	}
	return nil
}
