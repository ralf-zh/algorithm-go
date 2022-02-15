package nc0078

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
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
