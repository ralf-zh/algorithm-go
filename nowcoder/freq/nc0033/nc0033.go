package nc0033

/**
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here

	head, p1, p2 := choose(pHead1, pHead2)
	next := head

	for p1 != nil || p2 != nil {
		next.Next, p1, p2 = choose(p1, p2)
		next = next.Next
	}

	return head
}

func choose(p1 *ListNode, p2 *ListNode) (*ListNode, *ListNode, *ListNode) {
	if p1 == nil && p2 == nil {
		return nil, nil, nil
	}
	if p1 == nil && p2 != nil {
		return p2, p1, p2.Next
	}
	if p2 == nil && p1 != nil {
		return p1, p1.Next, p2
	}
	if p1.Val < p2.Val {
		return p1, p1.Next, p2
	}
	return p2, p1, p2.Next
}
