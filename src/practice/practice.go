package main

import (
	"fmt"
)

func main() {
	var intArry []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	aArry := intArry[0:len(intArry)]
	bArry := intArry[1:3]
	cArry := intArry[5:len(intArry)]
	dArry := intArry[5:]
	eArry := intArry[:]

	fmt.Printf("%+v\n", aArry)
	fmt.Printf("%+v\n", bArry)
	fmt.Printf("%+v\n", cArry)
	fmt.Printf("%+v\n", dArry)
	fmt.Printf("%+v\n", eArry)

}
