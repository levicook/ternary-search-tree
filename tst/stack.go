package tst

type stackElement struct {
	next  *stackElement
	value *node
}

type stack struct {
	head *stackElement
}

func (s *stack) pop() *node {
	el := s.head
	s.head = el.next
	return el.value
}

func (s *stack) push(n *node) {
	s.head = &stackElement{next: s.head, value: n}
}

func (s stack) empty() bool {
	return s.head == nil
}
