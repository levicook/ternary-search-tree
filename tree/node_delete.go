package tree

func (root *node) Delete(key string) {
	root.del([]rune(key))
}

func (n *node) del(runes []rune) {
	n = n.get(runes)

	if n != nil && n.End {
		n.End = false

		if n.isLeaf() { // then it can be deleted too
			switch n {
			case n.Par.Lo:
				n.Par.Lo = nil
			case n.Par.Eq:
				n.Par.Eq = nil
			case n.Par.Hi:
				n.Par.Hi = nil
			}
		}
	}
}
