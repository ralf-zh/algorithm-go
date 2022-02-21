package nc0061

func twoSum(nums []int, target int) []int {
	var exists = make(map[int]int, len(nums))

	for i, n := range nums {
		if another, ok := exists[target-n]; ok {
			return []int{another, i + 1}
		}
		exists[n] = i + 1
	}
	return nil
}
