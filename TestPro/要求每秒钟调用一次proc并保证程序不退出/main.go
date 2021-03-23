package main

import (
	"fmt"
	"time"
)

//package main
//
//func main() {
//    go func() {
//        // 1 在这里需要你写算法
//        // 2 要求每秒钟调用一次proc函数
//        // 3 要求程序不能退出
//    }()
//
//    select {}
//}
//
//func proc() {
//    panic("ok")
//}

func proc() {
	panic("ok")
}

func main() {
	go func() {
		for {
			timer := time.NewTimer(time.Second)
			select {
			case <- timer.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					proc()
				}()
				timer.Reset(time.Second)
			}
		}
	}()
	select {}
}
