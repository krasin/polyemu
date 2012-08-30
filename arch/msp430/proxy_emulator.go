package msp430

// Proxy emulator implements emu.Emulator by starting mspdebug.

import (
	"fmt"
	"os/exec"

	"github.com/krasin/polyemu/emu"
)

type proxyEmulator struct {
}

func NewProxyEmulator() (emu.Emulator, error) {
	if !checkMspDebug() {
		return nil, fmt.Errorf("Could not find mspdebug in PATH")
	}
	return new(proxyEmulator), nil
}

func checkMspDebug() bool {
	cmd := exec.Command("mspdebug", "--version")
	return cmd.Run() == nil
}

func (e *proxyEmulator) Step(st *emu.State, diff *emu.Diff) emu.Code {
	return emu.NotImplemented
}
