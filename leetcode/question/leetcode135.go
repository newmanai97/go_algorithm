package question

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	for i, _ := range candies {
		candies[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candies[i] < candies[i+1]+1 {
				candies[i] = candies[i+1] + 1
			}
		}
	}
	sum := 0
	for _, v := range candies {
		sum = sum + v
	}
	return sum
}
