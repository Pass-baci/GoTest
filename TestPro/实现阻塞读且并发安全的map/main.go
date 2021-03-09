package main

import (
	"sync"
	"time"
)

//GO里面MAP如何实现key不存在 get操作等待
//直到key存在或者超时，保证并发安全，且需要实现以下接口：

type Map struct {
	m map[string]interface{}
	lock sync.RWMutex
}

type sp interface {
	Out(key string, val interface{})
	Rd(key string, timeout time.Duration) interface{}
}

func main() {
}
