package bm05

import "sort"

func mergeKLists(lists []*ListNode) *ListNode {
	// write code here
	var head, tail *ListNode
	for len(lists) > 0 {
		sort.Slice(lists, func(i, j int) bool {
			if lists[i] == nil {
				return false
			}
			if lists[j] == nil {
				return true
			}
			return lists[i].Val < lists[j].Val
		})
		for i, l := range lists {
			if l == nil {
				lists = lists[:i]
				break
			}
		}
		if len(lists) == 0 || lists[0] == nil {
			break
		}
		if head == nil {
			head = lists[0]
			tail = head
		} else {
			tail.Next = lists[0]
			tail = tail.Next
		}
		lists[0] = lists[0].Next
	}
	return head
}
