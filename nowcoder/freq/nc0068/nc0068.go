package nc0068

/**
 *
 * @param number int整型
 * @return int整型
 */
func jumpFloor(number int) int {
	// write code here
	if number <= 2 {
		return number
	}
	i, j := 1, 2
	for k := 2; k < number; k++ {
		i, j = j, i+j
	}

	return j
}
