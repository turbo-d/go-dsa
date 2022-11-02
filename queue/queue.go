package queue

type Queue interface {
	Front() int
	Push(int)
	Pop() int
	Size() int
	Empty() bool
}

type SliceQueue []int

func (s SliceQueue) Front() int {
	if s == nil {
		return 0
	}

	return s[0]
}

func (s *SliceQueue) Push(val int) {
	*s = append(*s, val)
}

func (s *SliceQueue) Pop() int {
	val := (*s)[0]
	*s = (*s)[1:len(*s)]
	return val
}

func (s SliceQueue) Size() int {
	return len(s)
}

func (s SliceQueue) Empty() bool {
	if len(s) > 0 {
		return true
	}

	return false
}
