package main

import (
	"fmt"
	"sort"
	"strconv"
)

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	uniqueMap := make(map[string]int)
	sort.Ints(nums)
	fmt.Println("sorted nums: ")
	fmt.Println(nums)
	for i, _ := range nums {
		iNext := i + 1
		count := len(nums) - 1
		for iNext < count {
			fmt.Println("round %d, iNext %d, count %s", i, iNext, count)
			sum := nums[i] + nums[iNext] + nums[count]
			if sum == 0 {
				key := strconv.Itoa(nums[i]) + strconv.Itoa(nums[iNext]) + strconv.Itoa(nums[count])
				if _, ok := uniqueMap[key]; !ok {
					fmt.Println(key)
					uniqueMap[key] = 1
					result = append(result, []int{nums[i], nums[iNext], nums[count]})
				}
				iNext = iNext + 1
				count = count - 1
			} else if sum > 0 {
				count = count - 1
			} else if sum < 0 {
				iNext = iNext + 1
			}
		}
	}

	return result
}

func main() {
	args := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(args)
	fmt.Println(result)
}
