package dcpu16

import (
	"github.com/krasin/polyemu/emu"
	"testing"
)

var tests = []emu.Test{
	{
		Mem:     []byte{0x61, 0x88}, // SET X, 1
		WantReg: []uint64{RX: 1},
		N:       1,
	},
}

func TestSet(t *testing.T) {
	for _, tt := range tests {
		tt.Reg = make([]uint64, RegCount)
		e := new(Emulator)
		emu.RunTest(t, e, tt)
	}
}
