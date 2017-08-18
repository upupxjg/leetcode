package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 4}
	fmt.Printf("%v\n", twoSum0(nums, 6))
	fmt.Printf("%v\n", twoSum1(nums, 6))
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

func twoSum1(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		x := nums[i]
		if _, ok := m[x]; !ok {
			m[x] = i
		}
		y := target - x
		if x == y {
			continue
		}
		if j, ok := m[y]; ok {
			return []int{j, i}
		}
	}
	return nil
}
