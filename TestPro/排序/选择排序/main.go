package main

import (
	"fmt"
)

func selectsort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		max := nums[i]
		maxIndex := i
		for j := i+1; j < len(nums); j++ {
			if max < nums[j] {  //降序
				max = nums[j]
				maxIndex = j
			}
		}
		if maxIndex != i {
			nums[i], nums[maxIndex] = nums[maxIndex], nums[i]
		}
	}
	return nums
}

func main() {
	nums := []int{21,34,12,343,5461,23323,54543,54,1}
	ints := selectsort(nums)
	fmt.Println(ints)
	var slice1 []int
	fmt.Println(len(slice1))
	var slice=[]int{}
	fmt.Println(len(slice))
}
