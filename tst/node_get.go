package tst

func (root *node) Get(key string) (val interface{}) {
	n := root.get([]rune(key))

	if n != nil {
		val = n.val
	}

	return
}

func (n *node) get(runes []rune) *node {
	for _, key := range runes {
		switch {
		case key < n.key:
			n = n.lo
		case key > n.key:
			n = n.hi
		case key == n.key:
			n = n.eq
		}

		if n == nil {
			break
		}
	}

	lastKey := runes[len(runes)-1]
	if n != nil && n.key != lastKey {
		n = nil
	}

	return n
}
