package lc0001

func twoSum(nums []int, target int) []int {
	var exists = make(map[int]int, len(nums))

	for i, n := range nums {
		if another, ok := exists[target-n]; ok {
			return []int{another, i}
		}
		exists[n] = i
	}
	return nil
}
