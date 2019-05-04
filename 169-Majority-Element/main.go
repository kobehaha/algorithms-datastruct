package main

import (
	"fmt"
)

func majorityElement(nums []int) int {

	tempMap := make(map[int]int)
	majorityLenth := len(nums) / 2
	for _, item := range nums {
		if _, ok := tempMap[item]; !ok {
			tempMap[item] = 1
		} else {
			tempMap[item] = tempMap[item] + 1
		}
	}

	for key, value := range tempMap {
		if value > majorityLenth {
			return key
		}
	}

	return 1
}

func main() {
	input := []int{3, 2, 3}
	result := majorityElement(input)
	fmt.Println(result)
}
