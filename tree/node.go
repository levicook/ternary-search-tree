package tree

type node struct {
	Par *node

	Lo,
	Eq,
	Hi *node

	Key rune
	Val interface{}

	End bool
}
