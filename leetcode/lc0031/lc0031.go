package lc0031

import (
	"sort"
)

func nextPermutation(nums []int) {
	for n := len(nums) - 2; n >= 0; n-- {
		for j := len(nums) - 1; j > n; j-- {
			if nums[n] < nums[j] {
				nums[n], nums[j] = nums[j], nums[n]
				sort.Slice(nums[n+1:], func(i_, j_ int) bool {
					return nums[n+1+i_] < nums[n+1+j_]
				})
				return
			}
		}
	}
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
