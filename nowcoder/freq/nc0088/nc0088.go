package nc0088

func findKth(a []int, n int, K int) int {
	// write code here
	return quickSort(a, 0, n-1, n-K)
}

func quickSort(nums []int, low, high int, K int) int {
	p := patition(nums, low, high)
	if p == K {
		return nums[p]
	}
	if p < K {
		return quickSort(nums, p+1, high, K)
	}
	return quickSort(nums, low, p-1, K)
}

func patition(nums []int, low, high int) int {
	l, h := low, high
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
	return l
}
