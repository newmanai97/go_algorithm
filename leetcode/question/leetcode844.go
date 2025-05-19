package question

func backspaceCompare(s string, t string) bool {
	sarr := []byte(s)
	tarr := []byte(t)
	sskipnum, tskipnum := 0, 0
	i, j := len(sarr)-1, len(tarr)-1

	for true {
		for i >= 0 {
			if sarr[i] == '#' {
				sskipnum++
			} else {
				if sskipnum > 0 {
					sskipnum--
				} else {
					break
				}
			}
			i--
		}

		for j >= 0 {
			if tarr[j] == '#' {
				tskipnum++
			} else {
				if tskipnum > 0 {
					tskipnum--
				} else {
					break
				}
			}
			j--
		}
		if i < 0 || j < 0 {
			break
		}
		if sarr[i] != tarr[j] {
			return false
		}

		i--
		j--
	}
	if i == -1 && j == -1 {
		return true
	}
	return false
}
