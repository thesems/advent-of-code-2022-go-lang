package deque

import "testing"

func TestDequeInt(t *testing.T) {
	// Test with integers
	d := &Deque[int]{}

	if !d.IsEmpty() {
		t.Error("Expected deque to be empty")
	}

	d.Push(1)
	d.Push(2)
	d.Push(3)

	item, err := d.Front()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if *item != 1 {
		t.Errorf("Expected 1, got %d", *item)
	}

	item, err = d.Front()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if *item != 2 {
		t.Errorf("Expected 2, got %d", *item)
	}
}

func TestDequeString(t *testing.T) {
	// Test with strings
	dStr := &Deque[string]{}

	if !dStr.IsEmpty() {
		t.Error("Expected deque to be empty")
	}

	dStr.Push("A")
	dStr.Push("B")
	dStr.Push("C")

	itemStr, err := dStr.Front()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if *itemStr != "A" {
		t.Errorf("Expected 'A', got %s", *itemStr)
	}
}
