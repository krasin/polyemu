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
		Mem:     []byte{0x0f, 0x40}, // mov PC, r15
		WantReg: []uint64{R14: 2, PC: 2},
		N:       1,
	},
	{
		Mem:     []byte{0x0a, 0x41}, // mov SP, r10
		Reg:     []uint64{SP: 0x4000},
		WantReg: []uint64{SP: 0x40000, R10: 0x4000, PC: 2},
		N:       1,
	},
	{
		Mem:     []byte{0x0c, 0x42}, // mov r2, r12
		Reg:     []uint64{R2: 0xFEFE},
		WantReg: []uint64{R2: 0xFEFE, R12: 0xFEFE, PC: 2},
		N:       1,
	},
	{
		Mem:     []byte{0x0e, 0x43}, // mov r3, r14
		Reg:     []uint64{R4: 17},
		WantReg: []uint64{R4: 17, R14: 17, PC: 2},
		N:       1,
	},
	{
		Mem:     []byte{0x00, 0x4d}, // mov r13, r0
		Reg:     []uint64{R13: 50},
		WantReg: []uint64{R13: 50, PC: 50},
		N:       1,
	},
	{
		Mem:     []byte{0x01, 0x47}, // mov r7, r1
		Reg:     []uint64{R7: 15},
		WantReg: []uint64{R7: 15, SP: 15, PC: 2},
		N:       1,
	},
	{
		Mem:     []byte{0x02, 0x45}, // mov r5, r2
		Reg:     []uint64{R5: 5},
		WantReg: []uint64{R2: 5, R5: 5, PC: 2},
		N:       1,
	},
	{
		Mem:     []byte{0x03, 0x44}, // mov r4, r3
		Reg:     []uint64{R4: 10},
		WantReg: []uint64{R3: 10, R4: 10, PC: 2},
		N:       1,
	},
	{
		Mem:     []byte{0x0c, 0x4c}, // mov r12, r12
		Reg:     []uint64{R12: 77},
		WantReg: []uint64{R12: 77, PC: 2},
		N:       1,
	},
	{
		Mem:     []byte{0x00, 0x40}, // mov r0, r0
		WantReg: []uint64{PC: 2},
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
