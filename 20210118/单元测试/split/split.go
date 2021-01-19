package split

import (
	"fmt"
	"strings"
)

//应用于测试函数
func Split(s, sep string) (result []string) {
	//为Benchmark优化内存分配
	result = make([]string, 0, strings.Count(s, sep)+1)
	i := strings.Index(s, sep) //Index字符串查找，如果找不到则返回-1，字符串为空则返回0
	for i > -1 {
		result = append(result, s[:i])
		//s = s[i+1:] //TestSplit2出错
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return result
}

//用于测试覆盖率函数
func Split1() {
	fmt.Println("123")
}
