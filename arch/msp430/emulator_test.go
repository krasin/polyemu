package msp430

import (
	"github.com/krasin/polyemu/emu"
	"testing"
)

var tests = []emu.Test{}

func TestAll(t *testing.T) {
	for testInd, tt := range tests {
		tt.Reg = append(tt.Reg, make([]uint64, RegCount)...)
		tt.Mem = append(tt.Mem, make([]byte, 65536)...)
		e := new(Emulator)
		emu.RunTest(testInd, t, e, tt)
	}
}
