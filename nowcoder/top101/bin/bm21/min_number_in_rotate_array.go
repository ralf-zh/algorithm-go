package bm21

func minNumberInRotateArray(rotateArray []int) int {
	// write code here
	l, r := 0, len(rotateArray)-1
	for l < r {
		mid := l + (r-l)/2
		if rotateArray[mid] > rotateArray[r] {
			l = mid + 1
		} else if rotateArray[mid] < rotateArray[r] {
			r = mid
		} else {
			r--
		}
	}
	return rotateArray[l]
}
