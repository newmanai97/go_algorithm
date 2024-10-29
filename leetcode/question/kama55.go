package question

func rightReverseStr(str string, target int) string {
	strByte := []byte(str)

	reverse(strByte, 0, len(strByte)-1)
	reverse(strByte, 0, target-1)
	reverse(strByte, target, len(strByte)-1)

	return string(strByte)
}

func reverse(strByte []byte, i, j int) []byte {
	for i < j {
		strByte[i], strByte[j] = strByte[j], strByte[i]
		i++
		j--
	}
	return strByte
}

func rightReverseStr2(str string, target int) string {
	strByte := []byte(str)

	strByte = append(strByte[len(strByte)-target:], strByte...)
	strByte = strByte[:len(strByte)-target]
	return string(strByte)
}
