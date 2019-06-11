package main

import (
	"fmt"
	"regexp"
)

func main(){
	find()
	findall()
}

func find(){
	b := []byte("abc1def1")
	pat := `abc1|abc1def1`
	reg1 := regexp.MustCompile(pat)      // 第一匹配
	reg2 := regexp.MustCompilePOSIX(pat) // 最长匹配
	fmt.Printf("%s\n", reg1.Find(b))     // abc1
	fmt.Printf("%s\n", reg2.Find(b))     // abc1def1
}

func findall(){
	pat := `(((abc.)def.)ghi)`
	reg := regexp.MustCompile(pat)

	s := []byte(`abc-def-ghi abc+def+ghi`)

	// 查找所有匹配结果
	for _, one := range reg.FindAll(s, -1) {
		fmt.Printf("%s\n", one)
	}
	// abc-def-ghi
	// abc+def+ghi

	// 查找所有匹配结果及其分组字符串
	all := reg.FindAllSubmatch(s, -1)
	for i := 0; i < len(all); i++ {
		fmt.Println()
		one := all[i]
		for i := 0; i < len(one); i++ {
			fmt.Printf("%d: %s\n", i, one[i])
		}
	}
	// 0: abc-def-ghi
	// 1: abc-def-ghi
	// 2: abc-def-
	// 3: abc-

	// 0: abc+def+ghi
	// 1: abc+def+ghi
	// 2: abc+def+
	// 3: abc+
}