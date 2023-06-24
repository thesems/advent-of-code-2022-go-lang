package utils

import (
	"testing"
)

func TestStack1(t *testing.T) {
	stack := Stack{}
	stack.Push('A')
	stack.Push('B')
	stack.Push('C')
	stack.Pop()
	stack.Print()

	if len(stack.items) != 2 {
		t.Errorf("Stack items differ!")
	}
}

func TestStack2(t *testing.T) {
	stack := Stack{}
	stack.Push('A')
	stack.Push('B')
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Print()
}
