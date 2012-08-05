package dcpu16

import "testing"

func TestSetGet(t *testing.T) {
	mem := &memory{
		a:    make([]byte, 30),
		diff: make(map[uint64]byte),
	}
	mem.Set(10, 5)
	got := mem.At(10)
	if got != 5 {
		t.Errorf("Want: 5, got %d\n", got)
	}
}
