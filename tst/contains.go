package tst

func (root *node) Contains(key string) bool {
	return root.contains([]rune(key), 0)
}

func (n *node) contains(runes []rune, level int) bool {
	for {
		k := runes[level]

		switch {
		case k < n.key:
			n = n.lo
		case k > n.key:
			n = n.hi
		case k == n.key:
			n = n.eq
		}

		level++

		if n == nil || level == len(runes) {
			break
		}
	}

	return n != nil && n.end
}
