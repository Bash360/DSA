package stack

type Stack[T any] struct {
	data []T
	Size int
}

func New[T any](args ...T) *Stack[T] {
	return &Stack[T]{data: args, Size: len(args)}
}
func (self *Stack[T]) Push(val T) {
	self.data = append(self.data, val)
	self.Size = len(self.data)

}

func (self *Stack[T]) Pop() T {
	if self.Size == 0 {
		panic("empty stack")
	}

	result := self.data[self.Size-1]

	self.data = self.data[:self.Size-1]
	self.Size = len(self.data)
	return result

}

func (self *Stack[T]) Peek() T {
	if self.Size == 0 {
		panic("empty stack")

	}
	return self.data[self.Size-1]
}

func (self *Stack[T]) IsEmpty() bool {
	return len(self.data) == 0
}
