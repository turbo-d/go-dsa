package main

type Stack interface {
	Top() int
	Push(int)
	Pop() int
	Size() int
	Empty() bool
}

type SliceStack []int

func (s SliceStack) Top() int {
	if s != nil {
		return s[0]
	}

	return 0
}

func (s *SliceStack) Push(val int) {
	*s = append(*s, val)
}

func (s *SliceStack) Pop() int {
	n := len(*s)
	val := (*s)[n-1]
	*s = (*s)[:n-1]
	return val
}

func (s SliceStack) Size() int {
	return len(s)
}

func (s SliceStack) Empty() bool {
	return len(s) == 0
}
