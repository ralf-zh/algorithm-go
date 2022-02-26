package nc0004

/**
 *
 * @param head ListNode类
 * @return bool布尔型
 */
func hasCycle(head *ListNode) bool {
	// write code here
	if head == nil {
		return false
	}
	slow, fast := head, head.Next

	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	// write code here
	slow, fast := head, head

	for slow != nil && fast != nil {
		if fast.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}
