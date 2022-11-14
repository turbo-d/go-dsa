package deque

import "testing"

func TestFront(t *testing.T) {
	var s SliceDeque
	s.PushFront(4)
	s.PushFront(3)
	if front := s.Front(); front != 3 {
		t.Error("Expected 3 at the front of the deque, got ", front)
	}
	if size := s.Size(); size != 2 {
		t.Error("Expected a deque size of 2, got ", size)
	}
}

func TestPushFront(t *testing.T) {
	var s SliceDeque
	s.PushFront(3)
	if front := s.Front(); front != 3 {
		t.Error("Expected 3 at the front of the deque, got ", front)
	}
	if size := s.Size(); size != 1 {
		t.Error("Expected a deque size of 1, got ", size)
	}
}

func TestPopFront(t *testing.T) {
	var s SliceDeque
	s.PushFront(4)
	s.PushFront(3)
	if pop := s.PopFront(); pop != 3 {
		t.Error("Expected to pop 3 from the deque, got ", pop)
	}
	if size := s.Size(); size != 1 {
		t.Error("Expected a deque size of 1, got ", size)
	}
}

func TestBack(t *testing.T) {
	var s SliceDeque
	s.PushBack(4)
	s.PushBack(3)
	if back := s.Back(); back != 3 {
		t.Error("Expected 3 at the back of the deque, got ", back)
	}
	if size := s.Size(); size != 2 {
		t.Error("Expected a deque size of 2, got ", size)
	}
}

func TestPushBack(t *testing.T) {
	var s SliceDeque
	s.PushBack(3)
	if back := s.Back(); back != 3 {
		t.Error("Expected 3 at the back of the deque, got ", back)
	}
	if size := s.Size(); size != 1 {
		t.Error("Expected a deque size of 1, got ", size)
	}
}

func TestPopBack(t *testing.T) {
	var s SliceDeque
	s.PushBack(4)
	s.PushBack(3)
	if pop := s.PopBack(); pop != 3 {
		t.Error("Expected to pop 3 from the deque, got ", pop)
	}
	if size := s.Size(); size != 1 {
		t.Error("Expected a deque size of 1, got ", size)
	}
}

func TestSize(t *testing.T) {
	var s SliceDeque
	if size := s.Size(); size != 0 {
		t.Error("Expected a deque size of 0, got ", size)
	}
	s.PushFront(3)
	s.PushFront(4)
	if size := s.Size(); size != 2 {
		t.Error("Expected a deque size of 2, got ", size)
	}
}

func TestEmpty(t *testing.T) {
	var s SliceDeque
	if !s.Empty() {
		t.Error("Expected Empty() to return true, instead returned false")
	}
	s.PushFront(4)
	if s.Empty() {
		t.Error("Expected Empty() to return false, instead returned true")
	}
}
