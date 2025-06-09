package question

func lemonadeChange(bills []int) bool {
	five, ten, twenty := 0, 0, 0
	for _, v := range bills {
		if v == 5 {
			five++
		} else if v == 10 {
			if five <= 0 {
				return false
			}
			five--
			ten++
		} else if v == 20 {
			if five > 0 && ten > 0 {
				five--
				ten--
				twenty++
			} else if five >= 3 {
				five = five - 3
				twenty++
			} else {
				return false
			}
		}
	}
	return true
}
