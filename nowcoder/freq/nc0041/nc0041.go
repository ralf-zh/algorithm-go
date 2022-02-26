package nc0041

func maxLength(arr []int) int {
	// write code here
	var n2i = make(map[int]int, len(arr))

	// [3,3,2,1,3,3,3,1]
	var l = -1
	var max = 0
	for j, n := range arr {
		if i, ok := n2i[n]; ok {
			if i > l {
				l = i
			}
		}
		if j-l > max {
			max = j - l
		}

		n2i[n] = j
	}
	return max
}
