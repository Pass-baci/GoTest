package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{
		"root": 0,
		"user": 1,
		"guest": 2,
		"nobaby": 3,
	}
	var temp []string
	for key, _ := range m {
		temp = append(temp, key)
	}
	sort.Strings(temp)
	for _, v := range temp {
		fmt.Printf("key:%s, val:%d\n",v, m[v])
	}
}
