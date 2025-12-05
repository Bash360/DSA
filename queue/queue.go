package queue

import "fair-money/linkedlist"

type Queue[T any] struct {
	linkedlist *linkedlist.LinkedList[T]
	Size       int
}

func New[T any]() *Queue[T] {
	return &Queue[T]{linkedlist: linkedlist.New[T]()}
}

func (self *Queue[T]) Enqueue(val T) {
	self.linkedlist.InsertLast(val)
	self.Size = self.linkedlist.Size
}

func (self *Queue[T]) Dequeue() T {

	data := self.linkedlist.RemoveFirst()
	self.Size = self.linkedlist.Size
	return data
}

func (self *Queue[T]) Peek() T {
	return self.linkedlist.PeekFirst()
}

func (self *Queue[T]) IsEmpty() bool {
	return self.linkedlist.IsEmpty()
}
