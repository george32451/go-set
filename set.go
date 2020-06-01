package set

type Set struct {
	root *node
}

func (s *Set) Add(value string) error {
	if s.root == nil {
		s.root = &node{val: value}
		return nil
	}

	return s.root.Add(value)
}

func (s *Set) Size() int {
	if s.root == nil {
		return 0
	}
	return 1 + s.root.left.Size() + s.root.right.Size()
}
