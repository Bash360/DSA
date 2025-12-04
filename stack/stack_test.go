package stack

import "testing"

func TestInit(t *testing.T) {
	stack := New[int](1, 2, 3, 4)

	if stack.Size() != 4 && !stack.IsEmpty() {
		t.Error("TestStackInit: Init not working")
	}

}

func TestPop(t *testing.T) {

	stack := New[int](2, 5, 7, 8)
	got := stack.Pop()

	if got != 8 && stack.Size() != 3 {
		t.Error("TestStackPop: returns the last inserted data and removes it from the stack")
	}

}

func TestPush(t *testing.T) {
	stack := New[int]()
	stack.Push(12)
	if stack.IsEmpty() {
		t.Error("TestStackPop: error stack is empty ")
	}
}

func TestPopEmpty(t *testing.T) {
	stack := New[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Error("Test Pop empty not panicking")
		}
	}()

	stack.Pop()
}

func TestPeek(t *testing.T) {
	Stack := New[int](1, 2, 3, 4, 5)

	if Stack.Peek() != 5 {
		t.Error("Test peek error")
	}
}


func TestSize(t *testing.T){
Stack:=New[string]()

if Stack.Size()!=0{
	t.Error("Test Size method incorrect size")
}
}