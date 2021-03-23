package main

import (
	"context"
	"fmt"
	"time"
)

//假设有一个超长的切片，切片的元素类型为int，切片中的元素为乱序排列。
//限时5秒，使用多个goroutine查找切片中是否存在给定值，
//在找到目标值或者超时后立刻结束所有goroutine的执行。
//比如切片为：[23, 32, 78, 43, 76, 65, 345, 762, …… 915, 86]，查找的目标值为345，
//如果切片中存在目标值程序输出:"Found it!"并且立即取消仍在执行查找任务的goroutine。
//如果在超时时间未找到目标值程序输出:"Timeout! Not Found"，同时立即取消仍在执行查找任务的goroutine

func main() {
	s := []int{23, 32, 78, 43, 76, 65, 345, 762, 915, 86, 15, 5156, 516516,155}
	ctx, cancel := context.WithCancel(context.Background())
	timer := time.NewTimer(5 * time.Second)
	l := len(s)
	size := 3
	end := 0
	ch := make(chan bool)
	for i := 0; i < l; i+=size {
		end = i + size
		if end > l {
			end = l
		}
		go SelectSlice(ctx, ch, 23, s[i:end])
	}
	select {
	case <-ch:
		cancel()
		fmt.Println("Fount it!")
	case <-timer.C:
		cancel()
		fmt.Println("Time out!")
	}
}

func SelectSlice(ctx context.Context, ch chan bool, target int, s []int) {
	for _, v := range s {
		select {
		case <- ctx.Done():
			return
		default:
		}
		fmt.Println(v)
		if target == v {
			ch <- true
			return
		}
	}
}