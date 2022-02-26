package nc0050

/**
 *
 * @param head ListNode类
 * @param k int整型
 * @return ListNode类
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	head, tail, next := reverseList(head, k)
	for {
		if tail == nil {
			break
		}
		tail.Next, tail, next = reverseList(next, k)
	}
	return head
}

func reverseList(head *ListNode, k int) (*ListNode, *ListNode, *ListNode) {
	var (
		cur     = head
		newHead *ListNode
	)

	i := 0
	for ; cur != nil; cur = cur.Next {
		i++
	}
	if i < k {
		return head, nil, nil
	}
	cur = head

	for i := 0; i < k && cur != nil; i++ {
		tmp := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = tmp
	}
	return newHead, head, cur
}
