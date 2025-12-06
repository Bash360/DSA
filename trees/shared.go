package trees

import (
	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	Left  *Node[T]
	Right *Node[T]
	Data  T
}

type Tree[T constraints.Ordered] struct {
	Root *Node[T]
	Size int
}
