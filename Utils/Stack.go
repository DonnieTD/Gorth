package utils

type Stack struct {
	Vals []interface{}
}

func (s *Stack) Push(val interface{}) {
	s.Vals = append(s.Vals, val)
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(s.Vals) == 0 {
		return nil, false
	}
	top := s.Vals[len(s.Vals)-1]
	s.Vals = s.Vals[:len(s.Vals)-1]
	return top, true
}
