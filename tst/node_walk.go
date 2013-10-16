package tst

func (root *node) walk(limit int, callback func(*node)) {
	s := new(stack)
	s.push(root)

	for s.empty() {
		n := s.pop()

		callback(n)

		if n.lo != nil {
			s.push(n.lo)
		}

		if n.eq != nil {
			s.push(n.eq)
		}

		if n.hi != nil {
			s.push(n.hi)
		}
	}
}
