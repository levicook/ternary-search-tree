package tst

func (root *node) Contains(key string) bool {
	return root.contains([]rune(key))
}

func (n *node) contains(runes []rune) bool {
	n = n.get(runes)
	return n != nil && n.end
}
