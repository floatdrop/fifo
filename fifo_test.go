package fifo

import (
	"testing"
)

func TestFIFO(t *testing.T) {
	_ = New[int, int](128)
}