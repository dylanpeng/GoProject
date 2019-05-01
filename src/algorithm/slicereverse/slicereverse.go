package slicereverse

func ReverseSlice(s []int) []int {
	if s == nil || len(s) == 0 {
		return nil
	}
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return s
}
