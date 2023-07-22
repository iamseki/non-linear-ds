package binary_tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value int) {
	if n.Value < value {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	} else if n.Value > value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	}
}

func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}

	if n.Value < value {
		return n.Right.Search(value)
	} else if n.Value > value {
		return n.Left.Search(value)
	}
	return true
}

// TODO
func (n *Node) Print(node *Node) {

}

// TODO
func (n *Node) Height(node *Node) int {
	return 0
}

// TODO
func (n *Node) Depth(node *Node) int {
	return 0
}
