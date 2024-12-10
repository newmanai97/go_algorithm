package question

import "fmt"

func replaceNumber() {
	var strs []byte

	fmt.Scanln(&strs)

	for i := 0; i < len(strs); i++ {
		if strs[i] <= '9' && strs[i] >= '0' {
			insertElement := []byte{'n', 'u', 'm', 'b', 'e', 'r'}
			strs = append(strs[:i-1], append(insertElement, strs[i+1:]...)...)
			i = i + len(insertElement) - 1
		}
	}

	fmt.Println(string(strs))
}

func ReplaceNumber2() {
	var strs []byte

	fmt.Scanln(&strs)
	for i := 0; i < len(strs); i++ {
		if strs[i] <= '9' && strs[i] >= '0' {
			insertElement := []byte{'n', 'u', 'm', 'b', 'e', 'r'}
			strs = append(strs[:i-1], append(insertElement, strs[i+1:]...)...)
			i = i + len(insertElement) - 1
		}
	}

	fmt.Println(string(strs))
}
