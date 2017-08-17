package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 4}
	fmt.Printf("%v\n", twoSum0(nums, 6))
}
func twoSum0(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i != j && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
