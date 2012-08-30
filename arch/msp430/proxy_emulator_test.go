package msp430

import (
	"testing"

	"github.com/krasin/polyemu/emu"
)

func TestNewProxyEmulator(t *testing.T) {
	if _, err := NewProxyEmulator(); err != nil {
		t.Fatalf("Could not create ProxyEmulator. Consider installing mspdebug (for example, via sudo apt-get install mspdebug)")
	}
}

func TestProxyEmulator(t *testing.T) {
	for testInd, tt := range tests {
		tt.Reg = append(tt.Reg, make([]uint64, RegCount)...)
		tt.Mem = append(tt.Mem, make([]byte, 65536)...)
		e, err := NewProxyEmulator()
		if err != nil {
			t.Fatalf("NewProxyEmulator: %v", err)
		}
		emu.RunTest(testInd, t, e, tt)
	}
}
