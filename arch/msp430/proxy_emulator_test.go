package msp430

import (
	"testing"
)

func TestNewProxyEmulator(t *testing.T) {
	if _, err := NewProxyEmulator(); err != nil {
		t.Fatalf("Could not create ProxyEmulator. Consider installing mspdebug (for example, via sudo apt-get install mspdebug)")
	}
}
