package tst

type node struct {
	lo,
	eq,
	hi *node

	key rune
	val interface{}

	end bool
}
