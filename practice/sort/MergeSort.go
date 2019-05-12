package main

import "fmt"

func MergeSort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	mergeSort(arr, 0, arrLen-1)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {

	tmpArr := make([]int, end-start+1)

	i := start
	j := mid + 1

	k := 0

	for ; i <= mid && j <= end; k++ {
		if arr[i] < arr[j] {
			tmpArr[k] = arr[i]
			i = i + 1
		} else {
			tmpArr[k] = arr[j]
			j = j + 1
		}
	}

	for ; i <= mid; i++ {
		tmpArr[k] = arr[i]
		k++
	}

	for ; j <= end; j++ {
		tmpArr[k] = arr[j]
		k++
	}

	copy(arr[start:end+1], tmpArr)

}

func main() {

	arr := []int{5, 6, 10, 0, 3, 7}
	MergeSort(arr)
	fmt.Println(arr)

}
