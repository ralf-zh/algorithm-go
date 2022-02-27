package nc0053

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// write code here
	var (
		cur = head
		nth *ListNode
	)

	// 1→2→3→4→5, n = 2
	// 1→2, n = 2
	var it int
	for cur != nil {
		it++
		if it == n+1 {
			nth = head
		} else if it > n {
			nth = nth.Next
		}
		cur = cur.Next
	}
	if nth == nil {
		return head.Next
	}
	nth.Next = nth.Next.Next

	return head
}
