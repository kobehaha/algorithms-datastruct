package main

import "fmt"

func Compute(s1 string, s2 string) int {
	m := len(s1)
	n := len(s2)
	tempData := make([][]int, m+1)

	for i := 0; i < m+1; i++ {
		tempData[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s1[i-1] == s2[j-1] {
				tempData[i][j] = tempData[i-1][j-1] + 1
			}
		}
	}

	fmt.Println(tempData)

	longest := 0

	for i, _ := range tempData {
		for j, v := range tempData[i] {
			if longest < tempData[i][j] {
				longest = v
			}
		}
	}

	return longest

}

func main() {
	fmt.Println(Compute("test1", "test2"))
}
