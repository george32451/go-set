package set

import "errors"

type node struct {
	val   string
	left  *node
	right *node
}

func (n *node) Size() int {
	if n == nil {
		return 0
	}
	return 1 + n.left.Size() + n.right.Size()
}

func (n *node) Add(value string) error {

	if n == nil {
		return errors.New("tree is nil")
	}

	if n.val == value {
		return nil
	}

	if value < n.val {
		if n.left == nil {
			n.left = &node{val: value}
			return nil
		}
		return n.left.Add(value)
	}

	if value > n.val {
		if n.right == nil {
			n.right = &node{val: value}
			return nil
		}
		return n.right.Add(value)
	}
	return nil
}
