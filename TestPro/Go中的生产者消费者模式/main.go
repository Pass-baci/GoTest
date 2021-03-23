package main

import (
	"fmt"
	"strconv"
	"sync"
)

//数据模型
type Task struct {
	data string
}
var wg sync.WaitGroup

//生产者逻辑
func producer(tasks chan Task) {
	t := Task{}
	for i := 0; i < 10; i++ {
		t.data = strconv.Itoa(i)
		tasks <- t
	}
}

func producerDispatch(tasks chan Task) {
	defer close(tasks)
	producer(tasks)
}

//消费者逻辑
func consumer(task Task) {
	fmt.Printf("consum task:%v\n", task.data)
}

func consumerDispatch(tasks chan Task) {
	defer wg.Done()
	for task := range tasks {
		consumer(task)
	}
}

func main() {
	consumerNum := 10
	chLen := 50
	tasks := make(chan Task, chLen) //缓冲池
	go producerDispatch(tasks) //启动一个goroutine处理生产
	for i := 0; i < consumerNum; i++ {
		wg.Add(1)
		go consumerDispatch(tasks)
	}
	wg.Wait()
}
