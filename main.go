package main

import (
	"fmt"
	"strconv"
	"suanfa/leetcode/question"
)

func main() {
	array := []int{-4, -1, 0, 3, 10}
	result := question.SortedSquares2(array)
	for i, j := range result {
		fmt.Println(strconv.Itoa(i) + ":=" + strconv.Itoa(j))
	}
}
