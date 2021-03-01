package main

import "strings"

//请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。 给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。 保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】。
//
//解题思路
//
//这里有几个重点，第一个是ASCII字符，ASCII字符字符一共有256个，其中128个是常用字符，可以在键盘上输入。128之后的是键盘上无法找到的。
//
//然后是全部不同，也就是字符串中的字符没有重复的，再次，不准使用额外的储存结构，且字符串小于等于3000。

func isUniqueString(s string) bool {
	if strings.Count(s,"") > 3000 {
		return false
	}
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}

func main() {

}