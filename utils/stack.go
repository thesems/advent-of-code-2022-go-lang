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
	return ' ', errors.New("out of bounds")
}

func (s *Stack) Clear() {
	s.items = nil
}

func (s *Stack) GetTop() (rune, error) {
	if len(s.items) > 0 {
		return s.items[len(s.items)-1], nil
	}
	return ' ', errors.New("stack empty")
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
