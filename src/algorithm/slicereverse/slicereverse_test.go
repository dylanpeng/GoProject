package slicereverse

import (
	"fmt"
	"testing"
)

func TestReverseSlice(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(s)
	s = ReverseSlice(s)
	str := fmt.Sprint(s)
	fmt.Println(str)
	if str != "[6 5 4 3 2 1 0]" {
		t.Error("not reverse [6 5 4 3 2 1 0]")
	}
}
