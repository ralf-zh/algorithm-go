package bm06

import "testing"

func TestHasCircle(t *testing.T) {
	var cases = []struct {
		input  *ListNode
		k      int
		output bool
	}{
		{
			// {3,2,0,-4},1
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
			k:      1,
			output: true,
		},
		{
			// {1},-1
			input: &ListNode{
				Val: 1,
			},
			k:      -1,
			output: false,
		},
		{
			// {-1,-7,7,-4,19,6,-9,-5,-2,-5},6
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
			k:      6,
			output: true,
		},
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
			k:      -1,
			output: false,
		},
	}

	for i, c := range cases {
		var expected = hasCycle(makeCircle(c.input, c.k))
		if expected != c.output {
			t.Errorf("[%d] fails, %v != %v", i, expected, c.output)
		}
	}
}

func makeCircle(head *ListNode, k int) *ListNode {
	if k < 0 {
		return head
	}
	var (
		pivot *ListNode
		t     = head
		i     = 0
	)

	for t.Next != nil {
		if i == k {
			pivot = t
		}
		i++
		t = t.Next
	}
	t.Next = pivot

	return head
}
