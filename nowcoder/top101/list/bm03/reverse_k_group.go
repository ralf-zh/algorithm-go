package bm03

func reverseKGroup(head *ListNode, k int) *ListNode {
	// write code here
	if head == nil {
		return nil
	}

	var root, h, t *ListNode

	h = head
	for {
		if h == nil {
			break
		}
		h1, t1 := reverse(h, k)
		if root == nil {
			root = h1
		} else {
			t.Next = h1
		}
		t = h
		h = t1
	}
	return root
}

func reverse(head *ListNode, k int) (*ListNode, *ListNode) {
	var i int
	var pivot *ListNode
	for cur := head; cur != nil; cur = cur.Next {
		i++
		if i == k {
			pivot = cur
			break
		}
	}
	if pivot == nil {
		return head, nil
	}

	tail := pivot.Next
	pivot.Next = nil

	var h *ListNode
	for cur := head; cur != nil; {
		tmp := cur.Next
		cur.Next = h
		h = cur
		cur = tmp
	}
	return h, tail
}
