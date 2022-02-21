package nc0076

var stack1 []int
var stack2 []int

func Push(node int) {
	stack2 = append(stack2, node)
}

func Pop() int {
	if len(stack1) == 0 {
		for i := len(stack2) - 1; i >= 0; i-- {
			stack1 = append(stack1, stack2[i])
		}
		stack2 = nil
	}
	node := stack1[len(stack1)-1]
	stack1 = stack1[:len(stack1)-1]
	return node
}
