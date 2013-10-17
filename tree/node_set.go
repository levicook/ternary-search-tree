package tree

func (root *node) Set(key string, val interface{}) {
	if val == nil {
		val = key
	}
	root.set([]rune(key), 0, val)
}

func (n *node) set(runes []rune, level int, val interface{}) {
	var next *node

	switch key := runes[level]; true {

	case key < n.Key:
		if n.Lo == nil {
			n.Lo = new(node)
			n.Lo.Par = n
			n.Lo.Key = key
		}
		next = n.Lo

	case key > n.Key:
		if n.Hi == nil {
			n.Hi = new(node)
			n.Hi.Par = n
			n.Hi.Key = key
		}
		next = n.Hi

	case key == n.Key:
		if n.Eq == nil {
			n.Eq = new(node)
			n.Eq.Par = n
			n.Eq.Key = key
		}
		next = n.Eq

	}

	nextLevel := level + 1
	if len(runes) > nextLevel {
		next.set(runes, nextLevel, val)
	} else if len(runes) == nextLevel {
		next.End = true
		next.Val = val
	}
}
