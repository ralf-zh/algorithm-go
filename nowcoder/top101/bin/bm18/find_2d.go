package bm18

func Find(target int, array [][]int) bool {
	if len(array) == 0 || len(array[0]) == 0 {
		return false
	}
	// write code here
	i, j := len(array)-1, 0
	for i >= 0 && j < len(array[0]) {
		if target == array[i][j] {
			return true
		}
		if target > array[i][j] {
			j++
		} else {
			i--
		}
	}
	return false
}
