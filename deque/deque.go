package deque

type Deque interface {
	Front() int
	PushFront(int)
	PopFront() int
	Back() int
	PushBack(int)
	PopBack() int
	Size() int
	Empty() bool
}

type SliceDeque []int

func (s SliceDeque) Front() int {
	if s == nil {
		return 0
	}

	return s[0]
}

func (s *SliceDeque) PushFront(val int) {
	*s = append(SliceDeque{val}, (*s)...)
}

func (s *SliceDeque) PopFront() int {
	val := (*s)[0]
	*s = (*s)[1:len(*s)]
	return val
}

func (s SliceDeque) Back() int {
	if s == nil {
		return 0
	}

	return s[len(s)-1]
}

func (s *SliceDeque) PushBack(val int) {
	*s = append(*s, val)
}

func (s *SliceDeque) PopBack() int {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func (s SliceDeque) Size() int {
	return len(s)
}

func (s SliceDeque) Empty() bool {
	if len(s) > 0 {
		return false
	}

	return true
}
