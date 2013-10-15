package tst

func (root *node) AllWithPrefix(prefix string, limit int) []interface{} {
	return root.allWithPrefix([]rune(prefix), limit, 0)
}

func (root *node) allWithPrefix(runes []rune, limit, level int) (values []interface{}) {
	// find the deepest node, matching prefix
	// grab all nodes, with values, below that, until we hit our limit
	return
}
