package tree

func (root *node) Get(key string) (val interface{}) {
	n := root.get([]rune(key))

	if n != nil {
		val = n.Val
	}

	return
}

func (n *node) get(runes []rune) *node {
	for i, _ := range runes {
		key := runes[i]

		switch {
		case key < n.Key:
			n = n.Lo
		case key > n.Key:
			n = n.Hi
		case key == n.Key:
			n = n.Eq
		}

		if n == nil {
			break
		}
	}

	lastKey := runes[len(runes)-1]
	if n != nil && n.Key != lastKey {
		n = nil
	}

	return n
}
