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
	{
		Mem:     []byte{0x61, 0x8c, 0x64, 0x0c}, // SET X, 2 ; MUL X, X
		WantReg: []uint64{RX: 4},
		N:       2,
	},
}

func TestSet(t *testing.T) {
	for _, tt := range tests {
		tt.Reg = make([]uint64, RegCount)
		e := new(Emulator)
		emu.RunTest(t, e, tt)
	}
}
