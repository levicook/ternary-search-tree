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

	case key < n.key:
		if n.lo == nil {
			n.lo = new(node)
			n.lo.par = n
			n.lo.key = key
		}
		next = n.lo

	case key > n.key:
		if n.hi == nil {
			n.hi = new(node)
			n.hi.par = n
			n.hi.key = key
		}
		next = n.hi

	case key == n.key:
		if n.eq == nil {
			n.eq = new(node)
			n.eq.par = n
			n.eq.key = key
		}
		next = n.eq

	}

	nextLevel := level + 1
	if len(runes) > nextLevel {
		next.set(runes, nextLevel, val)
	} else if len(runes) == nextLevel {
		next.end = true
		next.val = val
	}
}
