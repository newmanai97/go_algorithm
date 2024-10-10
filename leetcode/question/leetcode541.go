package question

func ReverseStr(s string, k int) string {
	bytes := []byte(s)
	len := len(bytes)
	for i := 0; i < len; i = i + 2*k {
		p := i
		if i+k <= len-1 {
			p = i + k - 1
		} else {
			p = len - 1
		}
		for q := i; q < p; q, p = q+1, p-1 {
			bytes[q], bytes[p] = bytes[p], bytes[q]
		}
	}
	return string(bytes)
}
