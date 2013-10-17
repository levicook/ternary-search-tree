package tree

func (n *node) isLeaf() bool {
	return n.Lo == nil && n.Eq == nil && n.Hi == nil
}
