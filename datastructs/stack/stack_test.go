package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := MakeStack()
	if !s.IsEmpty() {
		t.Error("weird!")
	}

	Num := uint32(20)

	for i := uint32(0); i < Num; i++ {
		s.Push(i)

		if s.Len() != i+1 {
			t.Error("Push failure0")
		}
	}

	for i := uint32(0); i < Num; i++ {

		if val, _ := s.Peek(); val != Num-i-1 {
			t.Error("Peek failure0")
		}

		s.Pop()

		if s.Len() != Num-i-1 {
			t.Error("Pop failure0")
		}

	}

}
