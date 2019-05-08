package main

import "fmt"

func main() {
	arry := [][]int{{0, 1, 2, 3, 4}, {10, 11, 12, 13, 14}, {20, 21, 22, 23, 24}, {30, 31, 32, 33, 34}, {40, 41, 42, 43, 44}}
	// i, j, temp := 0, 0, -1
	for i, j := 0, 0; i < len(arry); i++ {
		for ; j < len(arry[i]); j++ {
			fmt.Printf("%d,", arry[i][j])
		}
		j = 0
		fmt.Println("")
	}

	CircleSort(arry, 0, len(arry)-1)

}

func CircleSort(arryslice [][]int, h, l int) {
	fmt.Println(h, l)
	if h > l {
		return
	}
	for j := h; j <= l; j++ {
		fmt.Printf("%d,", arryslice[h][j])
	}
	for i := h + 1; i < len(arryslice)-1-h; i++ {
		fmt.Printf("%d,", arryslice[i][l])
	}
	for j := l; j > h; j-- {
		fmt.Printf("%d,", arryslice[len(arryslice)-1-h][j])
	}
	for i := len(arryslice) - 1 - h; i > h; i-- {
		fmt.Printf("%d,", arryslice[i][h])
	}
	h++
	l--
	fmt.Println()
	CircleSort(arryslice, h, l)
}
