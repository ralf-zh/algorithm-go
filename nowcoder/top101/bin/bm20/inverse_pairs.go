package bm20

func InversePairs(data []int) int {
	// write code here
	if len(data) < 2 {
		return 0
	}
	p := pairs{nums: data}
	p.mergeSort(0, len(data)-1)
	return p.count
}

type pairs struct {
	nums  []int
	count int
}

func (p *pairs) mergeSort(left, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	p.mergeSort(left, mid)
	p.mergeSort(mid+1, right)
	p.merge(left, mid, right)
}

func (p *pairs) merge(left, mid, right int) {
	var sorted []int
	l, r := left, mid+1
	for l <= mid && r <= right {
		if p.nums[l] <= p.nums[r] {
			sorted = append(sorted, p.nums[l])
			l++
		} else {
			sorted = append(sorted, p.nums[r])
			r++
			p.count += mid + 1 - l
			p.count %= 1000000007
		}
	}
	for l <= mid {
		sorted = append(sorted, p.nums[l])
		l++
	}
	for r <= right {
		sorted = append(sorted, p.nums[r])
		r++
	}
	for i, n := range sorted {
		p.nums[left+i] = n
	}
}
