package tst

type Tree interface {
	AllWithPrefix(prefix string, limit int) (values []interface{})
	Contains(key string) bool
	Delete(key string)
	Get(key string) (val interface{})
	Set(key string, val interface{})
}

type node struct {
	loKid,
	eqKid,
	hiKid *node

	key rune
	val interface{}
}

func New() Tree {
	root := new(node)
	return root
}

func (n *node) AllWithPrefix(prefix string, limit int) (values []interface{}) { return }

func (root *node) Contains(key string) bool {
	return root.contains([]rune(key), 0)
}

func (n *node) contains(runes []rune, index int) bool {
	var (
		next *node
		key  = runes[index]
	)

	switch {
	case key < n.key:
		next = n.loKid
	case key > n.key:
		next = n.hiKid
	case key == n.key:
		next = n.eqKid
	}

	if next != nil {
		if len(runes) == index+1 && key == next.key {
			return true
		} else if len(runes) > index+1 { // recurse
			return next.contains(runes, index+1)
		}
	}

	return false
}

func (n *node) Delete(key string) { return }

func (n *node) Get(key string) (val interface{}) {
	return
}

func (root *node) Set(key string, val interface{}) {
	if val == nil {
		val = key
	}
	root.set([]rune(key), 0, val)
}

func (n *node) set(runes []rune, index int, val interface{}) {
	var next *node

	switch key := runes[index]; true {

	case key < n.key:
		if n.loKid == nil {
			n.loKid = new(node)
			n.loKid.key = key
		}
		next = n.loKid

	case key > n.key:
		if n.hiKid == nil {
			n.hiKid = new(node)
			n.hiKid.key = key
		}
		next = n.hiKid

	case key == n.key:
		if n.eqKid == nil {
			n.eqKid = new(node)
			n.eqKid.key = key
		}
		next = n.eqKid

	}

	if len(runes) == index+1 {
		// TODO tag node as end of word, right?

	} else if len(runes) > index+1 {
		next.set(runes, index+1, val)

	}
}
