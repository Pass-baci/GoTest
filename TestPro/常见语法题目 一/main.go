package main

import "fmt"

type Param map[string]interface{}

type Show struct {
	Param
}

func main() {
	s := new(Show)
	//s.Param["RMB"] = 10000  由于Param是定义为map类型，所以不能使用new进行初始化，需要使用make
	//panic: assignment to entry in nil map
	//修改如下：
	s.Param = make(Param)
	s.Param["RMB"] = 10000
	fmt.Println(s.Param["RMB"])
}
