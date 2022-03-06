package bm02

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// write code here
	if m == n {
		return head
	}

	var (
		left, right *ListNode
		i           int
	)

	for cur := head; cur != nil; cur = cur.Next {
		i++
		if i == m-1 {
			left = cur
		}
		if i == n {
			right = cur
		}
	}
	var h1, t1 *ListNode
	if left == nil {
		h1 = head
	} else {
		h1 = left.Next
	}
	t1 = right.Next
	right.Next = nil

	newHead, newTail := reverse(h1)
	if left == nil {
		head = newHead
	} else {
		left.Next = newHead
	}
	newTail.Next = t1

	return head
}

func reverse(root *ListNode) (*ListNode, *ListNode) {
	var (
		head *ListNode
		tail = root
	)
	for next := root; next != nil; {
		tmp := next.Next
		next.Next = head
		head = next
		next = tmp
	}
	return head, tail
}
