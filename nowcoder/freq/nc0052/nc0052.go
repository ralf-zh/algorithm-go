package nc0052

func isValid(s string) bool {
	// write code here
	var stk str
	for _, ch := range s {
		if stk.empty() {
			stk.push(ch)
			continue
		}
		if ch == '(' || ch == '{' || ch == '[' {
			stk.push(ch)
		} else {
			if stk.empty() {
				return false
			}
			if ch == ')' && stk.top() == '(' {
				stk.pop()
			}
			if ch == '}' && stk.top() == '{' {
				stk.pop()
			}
			if ch == ']' && stk.top() == '[' {
				stk.pop()
			}
		}
	}
	return stk.empty()
}

type str []rune

func (s *str) push(ch rune) {
	(*s) = append(*s, ch)
}

func (s *str) pop() rune {
	ch := s.top()
	*s = (*s)[:len(*s)-1]
	return ch
}

func (s *str) empty() bool {
	return s == nil || len(*s) == 0
}

func (s *str) top() rune {
	return (*s)[len(*s)-1]
}
