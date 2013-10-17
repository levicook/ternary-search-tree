package tree

func (root *node) Delete(key string) {
	root.del([]rune(key))
}

func (n *node) del(runes []rune) {
	n = n.get(runes)

	if n != nil && n.End {
		n.End = false

		if n.Lo == nil && n.Eq == nil && n.Hi == nil {
			// this node can be deleted too
			if n.Par.Lo == n {
				n.Par.Lo = nil
			} else if n.Par.Eq == n {
				n.Par.Eq = nil
			} else if n.Par.Hi == n {
				n.Par.Hi = nil
			}
		}
	}
}
