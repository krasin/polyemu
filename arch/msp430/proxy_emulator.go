package msp430

// Proxy emulator implements emu.Emulator by starting mspdebug.

import (
	"fmt"
	"io"
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
	s, err := newProxyState()
	if err != nil {
		return emu.InternalError
	}
	defer s.Close()
	if _, err = io.WriteString(s.stdin, "reset\n"+
		"regs\n"+
		"exit\n"); err != nil {
		return emu.InternalError
	}
	return emu.NotImplemented
}

type proxyState struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
}

func newProxyState() (s *proxyState, err error) {
	s = &proxyState{
		cmd: exec.Command("mspdebug", "sim"),
	}
	if s.stdin, err = s.cmd.StdinPipe(); err != nil {
		return
	}
	if s.stdout, err = s.cmd.StdoutPipe(); err != nil {
		return
	}
	err = s.cmd.Start()
	return
}

func (s *proxyState) Close() error {
	s.cmd.Process.Kill()
	return nil
}
