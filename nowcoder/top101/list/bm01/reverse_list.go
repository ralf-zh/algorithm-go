package bm01

func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	var head *ListNode
	for next := pHead; next != nil; {
		tmp := next.Next
		next.Next = head
		head = next
		next = tmp
	}
	return head
}
