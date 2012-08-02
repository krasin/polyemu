package dcpu16

import (
	"github.com/krasin/polyemu/emu"
	"testing"
)

var tests = []emu.Test{
	{
		Mem:     []byte{0x61, 0x88}, // SET X, 1
		WantReg: []uint64{RX: 1, PC: 1},
		N:       1,
	},
	{
		Mem:     []byte{0x61, 0x8c, 0x64, 0x0c}, // SET X, 2 ; MUL X, X
		WantReg: []uint64{RX: 4, PC: 2},
		N:       2,
	},
	{
		Mem:     []byte{0x81, 0xfc}, // SET Y, 30
		WantReg: []uint64{RY: 30, PC: 1},
		N:       1,
	},
	{
		Mem: []byte{
			0x81, 0xfc, // SET Y, 30
			0x81, 0x99, // SET [Y], 5
			0x61, 0x78, 0x1e, 0x00, // SET X, [30]
		},
		WantReg: []uint64{RX: 5, RY: 30, PC: 4},
		N:       3,
	},
}

func TestSet(t *testing.T) {
	for _, tt := range tests {
		tt.Reg = make([]uint64, RegCount)
		tt.Mem = append(tt.Mem, make([]byte, 65536)...)
		e := new(Emulator)
		emu.RunTest(t, e, tt)
	}
}
