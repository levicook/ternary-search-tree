package tree

type node struct {
	par *node

	lo,
	eq,
	hi *node

	key rune
	val interface{}

	end bool
}
