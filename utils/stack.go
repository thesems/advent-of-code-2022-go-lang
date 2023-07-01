package utils

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []rune
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (rune, error) {
	if s.IsEmpty() {
		return ' ', errors.New("out of bounds")
	}
	ch := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return ch, nil
}

func (s *Stack) Top() (rune, error) {
	if s.IsEmpty() {
		return ' ', errors.New("stack empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Print() {
	for i := len(s.items) - 1; i >= 0; i-- {
		symbol := ""
		if i == len(s.items)-1 {
			symbol = "<- top"
		} else if i == 0 {
			symbol = "<- bottom"
		}
		fmt.Printf("[%c] %s\n", s.items[i], symbol)
	}
}
