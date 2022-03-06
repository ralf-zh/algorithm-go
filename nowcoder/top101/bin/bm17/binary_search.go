package bm17

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	// write code here
	l, h := 0, len(nums)-1
	for l <= h {
		m := l + (h-l)/2
		println(l, h, m)
		if nums[m] == target {
			return m
		}
		if nums[m] < target {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return -1
}
