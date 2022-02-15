package nc0093

import "container/list"

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */
func LRU(operators [][]int, k int) []int {
	// write code here

	var r []int
	var cache = list.New()
	var fast = make(map[int]*list.Element, len(operators))
	for _, op := range operators {
		switch op[0] {
		case 1: // 插入
			fast[op[1]] = cache.PushFront([2]int{op[1], op[2]})
			if cache.Len() > k {
				kv := cache.Remove(cache.Back()).([2]int)
				delete(fast, kv[0])
			}

		case 2: // 查找
			if ele, ok := fast[op[1]]; ok {
				r = append(r, ele.Value.([2]int)[1])
				cache.MoveBefore(ele, cache.Front())
			} else {
				r = append(r, -1)
			}
		}
	}

	return r
}
