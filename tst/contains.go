package tst

func (root *node) Contains(key string) bool {
	return root.contains([]rune(key), 0)
}

func (n *node) contains(runes []rune, level int) bool {
	var (
		next *node
		key  = runes[level]
	)

	switch {
	case key < n.key:
		next = n.lo
	case key > n.key:
		next = n.hi
	case key == n.key:
		next = n.eq
	}

	if next != nil {
		nextLevel := level + 1

		if len(runes) == nextLevel && key == next.key {
			return next.end

		} else if len(runes) > nextLevel {
			return next.contains(runes, nextLevel)

		}
	}

	return false
}
