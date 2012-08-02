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
	{
		Mem: []byte{
			0x01, 0x98, // SET A, 5
			0x01, 0xa2, 0x01, 0x00, // SET [A+1], 7
			0x21, 0x78, 0x06, 0x00, // SET B, [6]
		},
		WantReg: []uint64{RA: 5, RB: 7, PC: 5},
		N:       3,
	},
	{
		Mem: []byte{
			0x01, 0xaf, // SET PUSH, 10
			0x61, 0x78, 0xFF, 0xFF, // SET X, [0xFFFF]
		},
		WantReg: []uint64{RX: 10, SP: 0xFFFF, PC: 3},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x88, // SET X, 1
			0x65, 0x80, // MLI X, -1
		},
		WantReg: []uint64{RX: 0xFFFF, EX: 0xFFFF, PC: 2},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0xa8, // SET X, 9
			0x66, 0x90, // DIV X, 3
		},
		WantReg: []uint64{RX: 3, PC: 2},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0xa8, // SET X, 9
			0x67, 0x7c, 0xfd, 0xff, // DVI X, -3
		},
		WantReg: []uint64{RX: 0xFFFD, PC: 3},
		N:       2,
	},
}

func TestSet(t *testing.T) {
	for _, tt := range tests {
		tt.Reg = make([]uint64, RegCount)
		tt.Mem = append(tt.Mem, make([]byte, 2*65536)...)
		e := new(Emulator)
		emu.RunTest(t, e, tt)
	}
}
