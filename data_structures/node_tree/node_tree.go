package node_tree

// see 2023/day08 for usage

type Network struct {
	Root *Node
}

type Node struct {
	Self  string
	Left  *Node
	Right *Node
}

func (n *Node) GoLeft() *Node {
	return n.Left
}

func (n *Node) GoRight() *Node {
	return n.Right
}

func (n *Node) IsEnd(end string) bool {
	return n.Self == end
}
