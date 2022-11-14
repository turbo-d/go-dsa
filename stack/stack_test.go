package stack

import (
	"testing"
)

func TestSliceStackTop(t *testing.T) {
	var s SliceStack
	s.Push(4)
	s.Push(3)
	if top := s.Top(); top != 3 {
		t.Error("Expected 3, got ", top)
	}
	if size := s.Size(); size != 2 {
		t.Error("Expected a stack size of 1, got ", size)
	}
}

func TestSliceStackPush(t *testing.T) {
	var s SliceStack
	s.Push(3)
	if top := s.Top(); top != 3 {
		t.Error("Expected 3 on the top of the stack, got ", top)
	}
	if size := s.Size(); size != 1 {
		t.Error("Expected a stack size of 1, got ", size)
	}
}

func TestSliceStackPop(t *testing.T) {
	var s SliceStack
	s.Push(3)
	if pop := s.Pop(); pop != 3 {
		t.Error("Expected to pop a 3 off of the stack, got ", pop)
	}
	if size := s.Size(); size != 0 {
		t.Error("Expected a stack size of 0, got ", size)
	}
}

func TestSliceStackSize(t *testing.T) {
	var s SliceStack
	s.Push(3)
	if size := s.Size(); size != 1 {
		t.Error("Expected a stack size of 1, got ", size)
	}
	s.Pop()
	if size := s.Size(); size != 0 {
		t.Error("Expected a stack size of 0, got ", size)
	}
}

func TestSliceStackEmpty(t *testing.T) {
	var s SliceStack
	s.Push(3)
	if s.Empty() {
		t.Error("Expected Empty() to return false, got ", true)
	}
	s.Pop()
	if !s.Empty() {
		t.Error("Expected Empty() to return true, got ", false)
	}
}
