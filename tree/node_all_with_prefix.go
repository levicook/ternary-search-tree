package tree

func (root *node) AllWithPrefix(prefix string, limit int) []interface{} {
	return root.allWithPrefix([]rune(prefix), limit)
}

func (root *node) allWithPrefix(runes []rune, limit int) (values []interface{}) {
	if n := root.get(runes); n != nil {
		n.walk(limit, func(child *node) {
			if child.end {
				values = append(values, child.val)
			}
		})
	}
	return
}
