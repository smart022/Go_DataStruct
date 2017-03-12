package stack

import "errors"

type vtype interface{}

type Stack struct {
	cot    []vtype
	length uint32
}

func MakeStack() *Stack {
	s := new(Stack)
	s.length = 0
	s.cot = make([]vtype, 0)
	return s
}

func (s *Stack) Len() uint32 {
	return s.length
}

func (s *Stack) IsEmpty() bool {
	return s.length == uint32(0)
}

// append to slice!!
func (s *Stack) Push(val vtype) {
	s.cot = append(s.cot, val)
	s.length++
}

func (s *Stack) Pop() (vtype, error) {
	if s.Len() == 0 {
		return nil, errors.New("Empty Stack!")
	}

	cur := s.length
	s.length--
	return s.cot[cur-1], nil
}

func (s *Stack) Peek() (vtype, error) {
	if s.Len() == 0 {
		return nil, errors.New("Empty Stack!")
	}
	return s.cot[s.Len()-1], nil
}
