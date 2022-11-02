package queue

import "testing"

func TestFront(t *testing.T) {
	var s SliceQueue
	s.Push(3)
	s.Push(4)
	if front := s.Front(); front != 3 {
		t.Error("Expected 3 at the front of the queue, got ", front)
	}
	if size := s.Size(); size != 2 {
		t.Error("Expected a queue size of 2, got ", size)
	}
}

func TestPush(t *testing.T) {
	var s SliceQueue
	s.Push(3)
	if front := s.Front(); front != 3 {
		t.Error("Expected 3 at the front of the queue, got ", front)
	}
	if size := s.Size(); size != 1 {
		t.Error("Expected a queue size of 1, got ", size)
	}
}

func TestPop(t *testing.T) {
	var s SliceQueue
	s.Push(3)
	s.Push(4)
	if pop := s.Pop(); pop != 3 {
		t.Error("Expected to pop 3 from the queue, got ", pop)
	}
	if size := s.Size(); size != 1 {
		t.Error("Expected a queue size of 1, got ", size)
	}
}

func TestSize(t *testing.T) {
	var s SliceQueue
	if size := s.Size(); size != 0 {
		t.Error("Expected a queue size of 0, got ", size)
	}
	s.Push(3)
	s.Push(4)
	if size := s.Size(); size != 2 {
		t.Error("Expected a queue size of 2, got ", size)
	}
}

func TestEmpty(t *testing.T) {
	var s SliceQueue
	if !s.Empty() {
		t.Error("Expected Empty() to return true, instead returned false")
	}
	s.Push(4)
	if s.Empty() {
		t.Error("Expected Empty() to return false, instead returned true")
	}
}
