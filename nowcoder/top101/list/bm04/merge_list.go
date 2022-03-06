package bm04

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil && pHead2 == nil {
		return nil
	}
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}
	var root, tail *ListNode

	h1, h2 := pHead1, pHead2
	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			if root == nil {
				root, tail = h1, h1
			} else {
				tail.Next = h1
				tail = tail.Next
			}
			h1 = h1.Next
		} else {
			if root == nil {
				root, tail = h2, h2
			} else {
				tail.Next = h2
				tail = tail.Next
			}
			h2 = h2.Next
		}
	}
	if h1 != nil {
		tail.Next = h1
	}
	if h2 != nil {
		tail.Next = h2
	}
	return root
}
