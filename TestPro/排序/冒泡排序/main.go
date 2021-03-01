package main

import "fmt"

func BubbleSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	for i := 0; i < len(nums)-1; i++ {
		isFlag := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isFlag = true
			}
		}
		if !isFlag {
			return nums
		}
	}
	return nums
}

func main() {
	nums := []int{1,4,6,8,9,23,2,2}
	sort := BubbleSort(nums)
	fmt.Println(sort)
}
