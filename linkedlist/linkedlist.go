package linkedlist

type Node[T any] struct {
	Data T
	Next *Node[T]
	Prev *Node[T]
}
type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	Size int
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (self *LinkedList[T]) InsertFirst(data T) {
	newNode := &Node[T]{Data: data, Next: self.head}
	if self.IsEmpty() {
		self.head = newNode
		self.tail = newNode
		self.Size++
		return
	}

	currentHead := self.head
	self.head = newNode
	currentHead.Prev = self.head
	self.Size++

}
func (self *LinkedList[T]) InsertLast(data T) {
	newNode := &Node[T]{Data: data, Prev: self.tail}
	if self.IsEmpty() {
		self.head = newNode
		self.tail = newNode
		self.Size++
		return
	}

	currentTail := self.tail
	self.tail = newNode
	currentTail.Next = self.tail
	self.Size++
}

func (self *LinkedList[T]) RemoveFirst() T {
	if self.IsEmpty() {
		panic("Empty Linked List")
	}
	head := self.head.Data
	self.head = self.head.Next
	if self.head != nil {
		self.head.Prev = nil
	} else {
		self.tail = nil
	}

	self.Size--

	return head
}

func (self *LinkedList[T]) RemoveLast() T {
	if self.IsEmpty() {
		panic("Empty Linked List")
	}
	tail := self.tail.Data

	self.tail = self.tail.Prev
	if self.tail != nil {
		self.tail.Next = nil
	} else {
		self.head = nil
	}

	self.Size--
	return tail
}

func (self *LinkedList[T]) PeekLast() T {

	if self.IsEmpty() {
		panic("Empty Linked List")
	}
	return self.tail.Data
}

func (self *LinkedList[T]) PeekFirst() T {
	if self.IsEmpty() {
		panic("Empty Linked List")
	}
	return self.head.Data
}

func (self *LinkedList[T]) IsEmpty() bool {
	return self.head == nil
}
