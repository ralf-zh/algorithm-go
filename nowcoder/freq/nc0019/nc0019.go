package nc0019

func FindGreatestSumOfSubArray(array []int) int {
	if len(array) == 0 {
		return 0
	}
	var sum = 0
	var max = array[0]

	for _, n := range array {
		if sum < 0 {
			sum = 0
		}
		sum += n
		if sum > max {
			max = sum
		}
	}
	return max
}
