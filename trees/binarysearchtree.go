package trees

import "golang.org/x/exp/constraints"

type BinarySearchTree[T constraints.Ordered] struct {
	Tree[T]
}

func NewBinaryST[T constraints.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

func (self *BinarySearchTree[T]) Insert(val T) {

	self.Root = self.insertNode(self.Root, val)
}

func (self *BinarySearchTree[T]) insertNode(node *Node[T], val T) *Node[T] {

	if node == nil {
		return self.insertNode(&Node[T]{Data: val}, val)
	}

	if val < node.Data {

		node.Left = self.insertNode(node.Left, val)

	}

	if val > node.Data {

		node.Right = self.insertNode(node.Right, val)

	}
	return node
}
