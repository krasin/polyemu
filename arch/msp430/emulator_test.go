package msp430

import (
	"github.com/krasin/polyemu/emu"
	"testing"
)

var tests = []emu.Test{
	{
		Mem:     []byte{0x06, 0x45}, // mov r5, r6
		Reg:     []uint64{R5: 1},
		WantReg: []uint64{R5: 1, R6: 1, PC: 2},
		N:       1,
	},
	{
		Mem:     []byte{0x08, 0x47}, // mov r7, r8
		Reg:     []uint64{R7: 5},
		WantReg: []uint64{R7: 5, R8: 5, PC: 2},
		N:       1,
	},
	{
		Mem:     []byte{0x0a, 0x49}, // mov r9, r10
		Reg:     []uint64{R9: 22},
		WantReg: []uint64{R9: 22, R10: 22, PC: 2},
		N:       1,
	},
	{
		Mem:     []byte{0x0c, 0x4b}, // mov r11, r12
		Reg:     []uint64{R11: 0xFFFF},
		WantReg: []uint64{R11: 0xFFFF, R12: 0xFFFF, PC: 2},
		N:       1,
	},
	{
		Mem:     []byte{0x0e, 0x4d}, // mov r13, r14
		Reg:     []uint64{R13: 0xFF},
		WantReg: []uint64{R13: 0xFF, R14: 0xFF, PC: 2},
		N:       1,
	},
}

func TestAll(t *testing.T) {
	for testInd, tt := range tests {
		tt.Reg = append(tt.Reg, make([]uint64, RegCount)...)
		tt.Mem = append(tt.Mem, make([]byte, 65536)...)
		e := new(Emulator)
		emu.RunTest(testInd, t, e, tt)
	}
}
