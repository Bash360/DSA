package stack

type Stack[T any] struct {
	Data []T
	Size int
}

func New[T any](args ...T) *Stack[T] {
	return &Stack[T]{Data: args, Size: len(args)}
}
func (self *Stack[T]) Push(val T) {
	self.Data = append(self.Data, val)
	self.Size = len(self.Data)

}

func (self *Stack[T]) Pop() T {
	if self.Size == 0 {
		panic("empty stack")
	}

	result := self.Data[self.Size-1]

	self.Data = self.Data[:self.Size-1]
	self.Size = len(self.Data)
	return result

}

func (self *Stack[T]) Peek() T {
	if self.Size == 0 {
		panic("empty stack")

	}
	return self.Data[self.Size-1]
}

func (self *Stack[T]) IsEmpty() bool {
	return len(self.Data) == 0
}
