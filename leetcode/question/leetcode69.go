package question

// 二分查找方法
func mySqrt(x int) int {
	left, right := 0, x

	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			if mid+1 <= x {
				if (mid+1)*(mid+1) > x {
					return mid
				} else {
					left = mid + 1
				}
			} else {
				return mid
			}
		} else if mid*mid > x {
			if mid-1 >= 0 {
				if (mid-1)*(mid-1) > x {
					return mid - 1
				} else {
					right = mid - 1
				}
			} else {
				return mid
			}
		}
	}
	return -1
}

func mySqrt2(x int) int {
	left, right := 0, x
	ans := -1

	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
