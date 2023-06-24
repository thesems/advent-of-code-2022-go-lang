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
	if len(s.items) > 0 {
		ch := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return ch, nil
	}
	return '-', errors.New("out of bounds")
}

func (s *Stack) Print() {
	for _, item := range s.items {
		fmt.Printf("[%c]\n", item)
	}
}
