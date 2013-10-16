package tst

func (root *node) AllWithPrefix(prefix string, limit int) []interface{} {
	return root.allWithPrefix([]rune(prefix), limit, 0)
}

func (n *node) allWithPrefix(runes []rune, limit, level int) (values []interface{}) {

	if n = n.get(runes); n != nil {
		n.walk(limit, func(child *node) {
			if child.end {
				values = append(values, child.val)
			}
		})
	}

	return
}
