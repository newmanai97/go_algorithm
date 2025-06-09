package question

func canCompleteCircuit(gas []int, cost []int) int {
	cursum, totalsum := 0, 0
	start := 0
	for i := 0; i < len(gas); i++ {
		cursum += gas[i] - cost[i]
		totalsum += gas[i] - cost[i]
		if cursum < 0 {
			start = i + 1
			cursum = 0
		}
	}
	if totalsum < 0 {
		return -1
	}
	return start
}
