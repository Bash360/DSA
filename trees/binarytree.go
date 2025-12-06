package trees

import (
	"DSA/queue"
	"DSA/stack"

	"golang.org/x/exp/constraints"
)

type BinaryTree[T constraints.Ordered] struct {
	Tree[T]
}

func NewBinaryTree[T constraints.Ordered]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}

func (self *BinaryTree[T]) FindNode(target T) *Node[T] {
	st := stack.New[*Node[T]]()
	st.Push(self.Root)

	for !st.IsEmpty() {
		current := st.Pop()
		if current.Data == target {
			return current
		}
		if current.Right != nil {
			st.Push(current.Right)
		}
		if current.Left != nil {
			st.Push(current.Left)
		}

	}
	return nil
}

func (self *BinaryTree[T]) FindNodeRecursion(current *Node[T], target T) *Node[T] {

	if current == nil {
		return nil
	}

	if current.Data == target {
		return current
	}

	result := self.FindNodeRecursion(current.Left, target)

	if result != nil {
		return result
	}

	result = self.FindNodeRecursion(current.Right, target)
	if result != nil {
		return result
	}

	return nil

}

func (self *BinaryTree[T]) BFS(target T) *Node[T] {
	q := queue.New[*Node[T]]()
	q.Enqueue(self.Root)

	for !q.IsEmpty() {

		current := q.Dequeue()
		if current.Data == target {
			return current
		}
		if current.Left != nil {
			q.Enqueue(current.Left)
		}

		if current.Right != nil {
			q.Enqueue(current.Right)
		}

	}
	return nil
}

func (self *BinaryTree[T]) Insert(target T) {
	newData := &Node[T]{Data: target}
	if self.Root == nil {
		self.Root = newData
		return
	}
	q := queue.New[*Node[T]]()
	q.Enqueue(self.Root)
	for q.IsEmpty() {
		current := q.Dequeue()
		if current.Left == nil {
			current.Left = newData
			return
		} else {
			q.Enqueue(current.Left)
		}

		if current.Right == nil {
			current.Right = newData
			return
		} else {
			q.Enqueue(current.Right)
		}

	}

}
