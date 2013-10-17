package tree

func (root *node) walk(limit int, callback func(*node)) {
	walked := 0

	s := new(stack)
	s.push(root)

	for {
		n := s.pop()

		if n.Lo != nil {
			s.push(n.Lo)
		}

		if n.Eq != nil {
			s.push(n.Eq)
		}

		if n.Hi != nil {
			s.push(n.Hi)
		}

		callback(n)
		walked++

		if s.empty() || limit > 0 && walked == limit {
			break
		}
	}
}
