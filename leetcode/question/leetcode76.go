package question

import "math"

func MinWindow(s string, t string) string {
	smap := make(map[uint8]int)
	tmap := make(map[uint8]int)

	for _, i := range t {
		tmap[uint8(i)]++
	}

	left, right := 0, 0

	lenth := math.MaxInt
	res := ""
	for right < len(s) {
		if tmap[s[right]] > 0 {
			smap[s[right]]++
		}
		for mapequal(smap, tmap) && left <= right {
			if lenth > right-left+1 {
				res = s[left : right+1]
				lenth = len(res)
			}
			if _, ok := tmap[s[left]]; ok {
				smap[s[left]]--
			}
			left++
		}
		right++
	}
	return res
}

func mapequal(smap, tmap map[uint8]int) bool {
	for k, v := range tmap {
		if smap[k] < v {
			return false
		}
	}
	return true
}
