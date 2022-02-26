package nc0004

import "testing"

func TestNC0004(t *testing.T) {
	var cases = []struct {
		input  *ListNode
		index  int
		output bool
	}{
		{
			input: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: -4,
						},
					},
				},
			},
			index:  1,
			output: true,
		},
		{
			input: &ListNode{
				Val: 1,
			},
			index:  -1,
			output: false,
		},
		{
			input: &ListNode{
				Val: -1,
				Next: &ListNode{
					Val: -7,
					Next: &ListNode{
						Val: 7,
						Next: &ListNode{
							Val: -4,
							Next: &ListNode{
								Val: 19,
								Next: &ListNode{
									Val: 6,
									Next: &ListNode{
										Val: -9,
										Next: &ListNode{
											Val: -5,
											Next: &ListNode{
												Val: -2,
												Next: &ListNode{
													Val: -5,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			index:  6,
			output: true,
		},
	}

	for i, c := range cases {
		tailToIndex(c.input, c.index)
		expected := hasCycle(c.input)
		if expected != c.output {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}

func tailToIndex(head *ListNode, idx int) {
	if idx < 0 {
		return
	}
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	for i := 0; head != nil; i++ {
		if i == idx {
			tail.Next = head
			break
		}
		head = head.Next
	}
}
