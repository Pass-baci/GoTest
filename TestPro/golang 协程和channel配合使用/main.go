package main

import (
	"fmt"
	"math/rand"
	"sync"
)

//写代码实现两个 goroutine，
//其中一个产生随机数并写入到 go channel 中，
//另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。

var wg sync.WaitGroup

func Send(ch chan int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		intn := rand.Intn(5)
		ch <- intn
	}
	close(ch)
}

func Get(ch chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go Send(ch)
	go Get(ch)
	wg.Wait()
}
