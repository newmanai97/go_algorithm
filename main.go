package main

import (
	"fmt"
	"suanfa/leetcode/question"
)

func main() {
	array := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	question.ReverseString(array)

	str := string(array[:])
	fmt.Println(str)
}
