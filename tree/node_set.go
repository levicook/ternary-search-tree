package tree

func (root *node) Set(key string, val interface{}) {
	if val == nil {
		val = key
	}
	root.set([]rune(key), val)
}

func (n *node) set(runes []rune, val interface{}) {
	for _, key := range runes {
		switch {
		case key < n.Key:
			if n.Lo == nil {
				n.Lo = new(node)
				n.Lo.Par = n
				n.Lo.Key = key
			}
			n = n.Lo
		case key > n.Key:
			if n.Hi == nil {
				n.Hi = new(node)
				n.Hi.Par = n
				n.Hi.Key = key
			}
			n = n.Hi
		case key == n.Key:
			if n.Eq == nil {
				n.Eq = new(node)
				n.Eq.Par = n
				n.Eq.Key = key
			}
			n = n.Eq
		}
	}
	n.End = true
	n.Val = val
}
