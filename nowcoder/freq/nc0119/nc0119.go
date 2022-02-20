package nc0119

func GetLeastNumbers_Solution(input []int, k int) []int {
	l := len(input)
	// write code here
	patition(input, 0, l-1, k)
	return input[:k]
}

func patition(nums []int, low, high int, k int) {
	if k <= low {
		return
	}
	l, h := low, high
	if l < h {
		pivot := nums[l]
		for l < h {
			for l < h && nums[h] >= pivot {
				h--
			}
			nums[l] = nums[h]

			for l < h && nums[l] <= pivot {
				l++
			}
			nums[h] = nums[l]
		}
		nums[l] = pivot
		patition(nums, low, l-1, k)
		patition(nums, l+1, high, k)
	}
}
