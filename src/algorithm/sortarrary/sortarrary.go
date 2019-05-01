package main

import (
	"fmt"
	"sync"
	"time"
)

type t struct {
	data int
}

func main() {
	var mu sync.WaitGroup
	mu.Add(1)
	defer mu.Wait()
	ch := make(chan int)
	go test(ch, &mu)
	for j := 0; j < 10; j++ {
		fmt.Println("send ", j)
		ch <- j
	}
	time.Sleep(time.Second * 10)

	fmt.Println("close ")
	close(ch)

	// fmt.Println(SortArrary([]int{10, 6, 2, 5, 8}))
}

func SortArrary(arrary []int) []int {
	for i := 0; i < len(arrary)-1; i++ {
		flag := false
		for j := 1; j < len(arrary)-i; j++ {
			if arrary[j-1] > arrary[j] {
				arrary[j-1], arrary[j] = arrary[j], arrary[j-1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return arrary
}

func test(c chan int, mu *sync.WaitGroup) {
	for t := range c {
		fmt.Println("get ", t)
	}
	fmt.Println("exit ")
	mu.Done()

}
