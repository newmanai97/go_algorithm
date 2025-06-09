package question

func partitionLabels(s string) []int {
	mymap := make(map[int32]int, 26)
	for i, v := range s {
		mymap[v-'a'] = i
	}

	ans := make([]int, 0)
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if right < mymap[int32(s[i]-'a')] {
			right = mymap[int32(s[i]-'a')]
		}
		if i == right {
			ans = append(ans, right-left+1)
			left = i + 1
		}
	}
	return ans
}
