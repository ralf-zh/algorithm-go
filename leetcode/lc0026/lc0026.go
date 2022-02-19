package lc0026

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	i, j := 0, 1
	for j < n {
		if nums[j] == nums[i] {
			j++
			continue
		}
		i++
		nums[i] = nums[j]
	}
	return i + 1
}
