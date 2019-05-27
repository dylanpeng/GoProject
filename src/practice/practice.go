package main

import (
	"fmt"
	"time"
)

type DriverMileageRank struct {
	ID       int64   `json:"id"`
	RealName string  `json:"real_name"`
	Rank     int     `json:"rank"`
	Mileage  float64 `json:"mileage"`
	Reward   float64 `json:"reward"`
}

func (e *DriverMileageRank) String() string {
	return fmt.Sprintf("%+v", *e)
}

func main() {
	rankDay := time.Now()
	timespan := time.Date(rankDay.Year(), rankDay.Month(), rankDay.Day(), 0, 0, 0, 0, time.Local).Unix()
	fmt.Println(timespan)
	fmt.Println(rankDay.Unix())
}

func InterceptString(input string) string {
	if input == "" {
		return ""
	}

	inputSlice := []rune(input)
	halfLen := int((len(inputSlice) - 1) / 2)

	for i := 0; i <= halfLen; i++ {
		inputSlice[i] = '*'
	}

	return string(inputSlice)
}

