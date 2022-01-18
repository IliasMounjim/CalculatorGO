package stack

import "errors"

// Stack is an abstract data type and is implemented as a slice of interface{}
// elements.
type Stack []interface{}

// New returns a new, empty stack.
func New() Stack {
	return make(Stack, 0)
}

// IsEmpty returns true if the stack s is empty.
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push pushes v onto the stack s.
func (s *Stack) Push(v interface{}) error {
	*s = append(*s, v)
	return nil
}

// Pop pops the top element from the stack s and returns it
func (s *Stack) Pop() (interface{}, error) {
    if len(*s) == 0 {
		return nil, errors.New("empty stack")
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, nil
}

// Top returns the top element of the stack s without popping it.
func (s *Stack) Top() (interface{}, error) {
	if len(*s) == 0 {
		return nil, errors.New("empty stack")
	}
	return (*s)[len(*s)-1], nil
}
