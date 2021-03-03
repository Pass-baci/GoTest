package main

import (
	"fmt"
	"strings"
)

//给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
//这里规定【大小写为不同字符】，且考虑字符串重点空格。给定一个string s1和一个string s2，请返回一个bool，代表两串是否重新排列后可相同。
//保证两串的长度都小于等于5000。

func isRegroup(s1, s2 string) bool {
	s1l := len([]rune(s1))
	s2l := len([]rune(s1))
	if s1l > 5000 || s2l > 5000 || s1l != s2l {
		return false
	}
	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}
func main() {
	s1 := "hello 你好!"
	s2 := "你好 !hello"
	regroup := isRegroup(s1, s2)
	fmt.Println(regroup)
}

