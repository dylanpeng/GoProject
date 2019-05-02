package main

import (
	"fmt"
	"strconv"
)

func main() {
	ipstr, i := "192.168.0.1", 0
	var ipslice []string

	for _, data := range ipstr {
		d := string(data)
		if d != "." {
			if len(ipslice) < i+1 {
				ipslice = append(ipslice, d)
			} else {
				ipslice[i] += d
			}
		} else {
			i++
		}
	}

	var result int
	for index, str := range ipslice {
		num, _ := strconv.ParseInt(str, 10, 32)
		result += int(num) << uint(8*(len(ipslice)-index-1))
		fmt.Println(str)
	}
	fmt.Println(result)
}
