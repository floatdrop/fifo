package fifo

import (
	"testing"
)

func TestFIFO_zero(t *testing.T) {
	f := New[int, int](0)

	if e := f.Push(1, 1); e == nil || e.Value != 1 {
		t.Fatalf("first element should be evicted")
	}

	if e := f.Get(1); e != nil {
		t.Fatalf("fifo must be empty")
	}
}

func TestFIFO(t *testing.T) {
	f := New[int, int](2)

	if f.Len() != 0 {
		t.Fatalf("fifo must be empty")
	}

	if f.Push(1, 1) != nil {
		t.Fatalf("first element should not be evicted")
	}

	if f.Push(2, 2) != nil {
		t.Fatalf("second element should not be evicted")
	}

	if e := f.Push(3, 3); e == nil || e.Value != 1 {
		t.Fatalf("third element should cause eviction of first")
	}

	if f.Remove(1) != nil {
		t.Fatalf("first element should already be evicted")
	}

	if e := f.Remove(2); e == nil || *e != 2 {
		t.Fatalf("second element should be removed")
	}

	if e := f.Push(4, 4); e != nil {
		t.Fatalf("forth element should not cause eviction")
	}

	if f.Get(3) == nil {
		t.Fatalf("third element should be in fifo")
	}

	if e := f.Push(5, 5); e == nil || e.Key != 3 || e.Value != 3 {
		t.Fatalf("third element not be rearranged after get: %+v", e)
	}

	if e := f.Push(4, 8); e == nil || e.Key != 4 || e.Value != 4 {
		t.Fatalf("replacing with push should evict old value: %+v", e)
	}

	if e := f.Push(6, 6); e == nil || e.Value != 5 {
		t.Fatalf("updated key should be moved in front of queue: %+v", e)
	}
}