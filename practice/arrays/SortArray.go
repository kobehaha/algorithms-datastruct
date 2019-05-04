package main

import (
	"fmt"
)

type SortArray struct {
	Data  []int
	Count int
	Size  int
}

func SortArrayInit(capacity int) *SortArray {
	if capacity == 0 {
		return nil
	}
	return &SortArray{
		Data:  make([]int, capacity, capacity),
		Count: 0,
		Size:  capacity,
	}

}

func (sortArray *SortArray) Insert(item int) bool {

	// full
	if sortArray.Count == sortArray.Size {
		return false
	}

	tempI := 0

	for tempI = 0; tempI < sortArray.Count; tempI++ {
		if item < sortArray.Data[tempI] {
			break
		}

	}
	for j := sortArray.Count; j > tempI; j-- {
		sortArray.Data[j] = sortArray.Data[j-1]
	}
	sortArray.Data[tempI] = item
	sortArray.Count = sortArray.Count + 1

	return true

}

func (sortArray *SortArray) Find(item int) (result bool, value int) {

	if sortArray.Count == 0 {
		return false, 0
	}
	low := 0
	high := sortArray.Count - 1
	mid := 0

	for low <= high {

		mid = (low + high) / 2
		if sortArray.Data[mid] > item {
			high = mid - 1
		} else if sortArray.Data[mid] < item {
			low = mid + 1
		} else {
			return true, mid
		}
	}
	return false, 0

}

func (sortArray *SortArray) Delete(item int) bool {
	if sortArray.Count == 0 {
		return false
	}
	find, mid := sortArray.Find(item)
	if find {
		for i := mid; i < sortArray.Count; i++ {
			sortArray.Data[i] = sortArray.Data[i+1]
		}
		sortArray.Count = sortArray.Count - 1
		return true
	}
	return false
}

func main() {
	testArray := SortArrayInit(20)
	fmt.Println(testArray.Insert(2))
	fmt.Println(testArray.Insert(50))
	fmt.Println(testArray.Insert(150))
	fmt.Println(testArray.Insert(-1))
	fmt.Println(testArray.Insert(60))
	fmt.Println(testArray.Insert(20))
	fmt.Println(testArray.Data)
	// find
	fmt.Println(testArray.Find(150))
	fmt.Println(testArray.Delete(150))
	fmt.Println(testArray.Data)
}
