package set

import "errors"

type Node struct {
	val   string
	left  *Node
	right *Node
}

func (n *Node) Size() int {
	if n == nil {
		return 0
	}
	return 1 + n.left.Size() + n.right.Size()
}

func (n *Node) Add(value string) error {

	if n == nil {
		return errors.New("tree is nil")
	}

	if n.val == value {
		return nil
	}

	if value < n.val {
		if n.left == nil {
			n.left = &Node{val: value}
			return nil
		}
		return n.left.Add(value)
	}

	if value > n.val {
		if n.right == nil {
			n.right = &Node{val: value}
			return nil
		}
		return n.right.Add(value)
	}
	return nil
}
