package main

import (
	"fmt"
	"sort"
)

func firstMissingPostive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	sort.Ints(nums)
	index := len(nums) - 1
	if nums[index] <= 0 {
		return 1
	}
	if nums[0] > 1 {
		return 1
	}
	for i, v := range nums {
		if i < index {
			if v == 0 && nums[i+1] > 1 {
				return 1
			}
			if v < 0 && nums[i+1] > 1 {
				return 1
			}
			if v > 0 && nums[i+1] == v {
				continue
			}
			if v > 0 && nums[i+1] != v+1 {
				return v + 1
			}
		} else {
			return v + 1
		}

	}
	fmt.Println("---ddd")
	return 1
}

func main() {
	input := []int{0, 3}
	result := firstMissingPostive(input)
	fmt.Println(result)
}
