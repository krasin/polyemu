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
}

func TestAll(t *testing.T) {
	for testInd, tt := range tests {
		tt.Reg = append(tt.Reg, make([]uint64, RegCount)...)
		tt.Mem = append(tt.Mem, make([]byte, 65536)...)
		e := new(Emulator)
		emu.RunTest(testInd, t, e, tt)
	}
}
