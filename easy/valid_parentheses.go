package easy

func ValidParentheses(chars string) bool {
	closedBrackets := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	s := Stack{}
	for _, c := range chars {
		b, ok := closedBrackets[string(c)]
		if ok {
			if v := s.Pop(); v == nil || v != b {
				return false
			}
			continue
		}
		s.Push(string(c))
	}

	return s.len == 0
}

type Stack struct {
	Head *Node
	len  int
}

type Node struct {
	Value any
	next  *Node
}

func (s *Stack) Push(value any) {
	s.Head = &Node{
		Value: value,
		next:  s.Head,
	}
	s.len++
}

func (s *Stack) Pop() any {
	if s.len == 0 {
		return nil
	}
	v := s.Head.Value
	s.Head, s.Head.next = s.Head.next, nil
	s.len--
	return v
}

func (s *Stack) Len() int {
	return s.len
}
