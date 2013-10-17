package tree

func (root *node) Delete(key string) {
	root.del([]rune(key))
}

func (n *node) del(runes []rune) {
	n = n.get(runes)

	if n != nil && n.end {
		n.end = false

		if n.lo == nil && n.eq == nil && n.hi == nil {
			// this node can be deleted too
			if n.par.lo == n {
				n.par.lo = nil
			} else if n.par.eq == n {
				n.par.eq = nil
			} else if n.par.hi == n {
				n.par.hi = nil
			}
		}
	}
}
