package bm06

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
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
