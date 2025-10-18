package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	i := 1
	for j := 1; j < n; j++ {
		if nums[j] != nums[j-1] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	var nums = []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}
