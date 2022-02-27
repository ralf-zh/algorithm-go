package nc0003

import (
	"strconv"
	"testing"
)

func TestNC0003(t *testing.T) {
	var cases = []struct {
		input1 []int
		input2 []int
		output string
	}{
		{
			input1: []int{1, 2},
			input2: []int{3, 4, 5},
			output: "3",
		},
		{
			input1: []int{1},
			input2: nil,
			output: "null",
		},
		{
			input1: nil,
			input2: []int{2},
			output: "2",
		},
	}
	for i, c := range cases {
		l := EntryNodeOfLoop(makeList(c.input1, c.input2))
		var expeced = "null"
		if l != nil {
			expeced = strconv.Itoa(l.Val)
		}
		if expeced != c.output {
			t.Errorf("[%d] fails, %v != %v", i, expeced, c.output)
		}
	}
}

func makeList(input1 []int, input2 []int) *ListNode {
	h1, t1 := buildList(input1, false)
	h2, _ := buildList(input2, true)

	if h1 != nil {
		t1.Next = h2
	} else {
		h1 = h2
	}
	return h1
}

func buildList(input []int, cyclic bool) (*ListNode, *ListNode) {
	var head, tail *ListNode
	for _, n := range input {
		if head == nil {
			head = &ListNode{
				Val: n,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val: n,
			}
			tail = tail.Next
		}
	}
	if cyclic && tail != nil {
		tail.Next = head
	}
	return head, tail
}
