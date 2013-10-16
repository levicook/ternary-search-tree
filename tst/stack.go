package tst

type stackElement struct {
	next *stackElement
	node *node
}

type stack struct {
	head *stackElement
}

func (s *stack) pop() *node {
	el := s.head
	s.head = el.next
	return el.node
}

func (s *stack) push(n *node) {
	s.head = &stackElement{next: s.head, node: n}
}

func (s stack) empty() bool {
	return s.head == nil
}
