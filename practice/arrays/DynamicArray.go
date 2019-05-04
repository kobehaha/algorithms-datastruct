package main

import (
	"fmt"
)

type DynamicArray struct {
	Data  []int
	Count int
	Size  int
}

func DynamicArrayInit(size int) *DynamicArray {
	data := make([]int, size)
	array := &DynamicArray{
		Data:  data,
		Count: 0,
		Size:  size,
	}
	return array
}

func (array *DynamicArray) ExpandArray(arrayValues []int) []int {
	data := make([]int, len(arrayValues)*2)
	fmt.Println("--- start expand array ---")
	copy(data, arrayValues)
	return data
}

func (array *DynamicArray) Append(item int) bool {

	if array.Count < array.Size-1 {
		array.Data[array.Count+1] = item
	} else {
		arrayData := array.ExpandArray(array.Data)
		array.Data = arrayData
		array.Size = array.Size * 2
	}
	array.Count = array.Count + 1
	return true
}

func main() {
	arrayTest := DynamicArrayInit(5)
	fmt.Println("start ---> test", arrayTest)
	for i := 0; i < 100000; i++ {
		arrayTest.Append(i)
	}

}
