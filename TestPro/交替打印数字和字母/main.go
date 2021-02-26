package main

import (
	"fmt"
	"sync"
	"time"
)

//使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母
//最终效果如下：
//12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
var (
	wg       sync.WaitGroup
	letterCh = make(chan bool)
	numberCh = make(chan bool)
)

func main() {
	wg.Add(1)
	go number(numberCh)
	wg.Add(1)
	go letter(letterCh)
	numberCh <- true
	wg.Wait()
}

func number(numberCh chan bool) {
	i := 1
	for {
		select {
		case <-numberCh:
			fmt.Print(i)
			fmt.Print(i + 1)
			i += 2
			letterCh <- true
			break
		case <-time.After(1 * time.Second):
			fmt.Println("exit")
			close(numberCh)
			wg.Done()
			return
		}
	}
}

func letter(letterCh chan bool) {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for {
		select {
		case <-letterCh:
			if str == "" {
				close(letterCh)
				wg.Done()
				return
			}
			fmt.Print(str[0:1])
			fmt.Print(str[1:2])
			if len(str) == 0 {
				str = ""
			} else {
				str = str[2:]
			}
			numberCh <- true
			break
		default:
			break
		}
	}
}
