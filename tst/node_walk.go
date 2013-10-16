package tst

func (root *node) walk(limit int, callback func(*node)) {
	walked := 0

	s := new(stack)
	s.push(root)

	for s.empty() {
		n := s.pop()

		if n.lo != nil {
			s.push(n.lo)
		}

		if n.eq != nil {
			s.push(n.eq)
		}

		if n.hi != nil {
			s.push(n.hi)
		}

		callback(n)
		walked++

		if limit > 0 && walked == limit {
			break
		}
	}
}
