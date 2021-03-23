package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//场景：在一个高并发的web服务器中，要限制IP的频繁访问。
//现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。
//每个IP三分钟之内只能访问一次。修改以下代码完成该过程，要求能成功输出 success:100

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//type Ban struct {
//	visitIPs map[string]time.Time
//}
//
//func NewBan() *Ban {
//	return &Ban{visitIPs: make(map[string]time.Time)}
//}
//func (o *Ban) visit(ip string) bool {
//	if _, ok := o.visitIPs[ip]; ok {
//		return true
//	}
//	o.visitIPs[ip] = time.Now()
//	return false
//}
//func main() {
//	success := 0
//	ban := NewBan()
//	for i := 0; i < 1000; i++ {
//		for j := 0; j < 100; j++ {
//			go func() {
//				ip := fmt.Sprintf("192.168.1.%d", j)
//				if !ban.visit(ip) {
//					success++
//				}
//			}()
//		}
//
//	}
//	fmt.Println("success:", success)
//}

type Ban struct {
	visitIPs map[string]time.Time
	RW sync.RWMutex
}

func NewBan(ch chan bool) *Ban {
	o := &Ban{visitIPs: make(map[string]time.Time)}
	//启动一个goroutine进行监听操作，每过3分钟进行遍历一次visitIPs，从而得知时间访问时间间隔是否超过三分钟
	go func() {
		//创建一个定时器
		timer := time.NewTimer(time.Minute * 3)
		for {
			select {
			//case <-ctx.Done():
			//	return
			case <-ch:  //使用channel进行通知退出goroutine
				return
			case <-timer.C:
				for ip, oldtime := range o.visitIPs {
					//条件逻辑，进行判断间隔是否超过3分钟
					if time.Now().Sub(oldtime) >= time.Minute*3 {
						delete(o.visitIPs, ip)
					}
				}
				//重置定时时间
				timer.Reset(time.Minute * 3)
			}
		}
	}()
	return o
}
func (o *Ban) visit(ip string) bool {
	//因为map并不是并发安全的，在多个goroutine操作同一个map时，需要加锁
	o.RW.Lock()
	defer o.RW.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}
func main() {
	success := int64(0)
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//ban := NewBan(ctx)
	ch := make(chan bool)
	defer close(ch)
	ban := NewBan(ch)
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					//在多goroutine下操作同一个int值，会产生不同步情况，所以使用原子操作
					atomic.AddInt64(&success, 1)
				}
			}(j) //启动groutine时需要将参数传进去，不能使用外面的变量，否则，当函数执行时，外面的变量在变化，使得结果不一致
		}
	}
	wg.Wait()
	fmt.Println("success:", success)
}
