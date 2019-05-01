package main

import (
	"fmt"
)

func main() {
	fmt.Println(Fibonacci(3))
	fmt.Println(Fibonacci(4))
	fmt.Println(Fibonacci(5))
	fmt.Println(Fibonacci(6))
	fmt.Println(GetFibonacci(7))
	fmt.Println(GetFibonacci(8))
	fmt.Println(GetFibonacci(9))

}

func Fibonacci(num int) int {
	if num < 1 {
		return 0
	} else if num == 1 || num == 2 {
		return 1
	}
	return Fibonacci(num-2) + Fibonacci(num-1)
}

func GetFibonacci(num int) int {
	if num < 1 {
		return 0
	}
	f := Fibonaccigo()
	result := 0
	for i := 0; i < num; i++ {
		result = f()
	}
	return result
}

func Fibonaccigo() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
