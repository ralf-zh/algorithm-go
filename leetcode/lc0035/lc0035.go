package lc0035

func searchInsert(nums []int, target int) int {
	var mid int
	i, j := 0, len(nums)-1
	for i <= j {
		mid = (i + j) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return i
}
