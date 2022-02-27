package nc0003

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	var (
		slow, fast = pHead, pHead
	)
	var hasCircle bool
	for slow != nil && fast != nil {
		if fast.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			hasCircle = true
			break
		}
	}
	if !hasCircle {
		return nil
	}
	fast = pHead
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
