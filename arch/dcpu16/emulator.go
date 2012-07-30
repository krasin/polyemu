package dcpu16

import (
	"fmt"

	"github.com/krasin/polyemu/emu"
)

const (
	PC       = iota
	RegCount = iota
)

type Emulator struct {
}

type regState []uint64

func (r regState) PC() uint16 {
	return uint16(r[PC])
}

type state struct {
	mem []byte
	reg regState
}

func (st *state) Fetch() emu.Code {
	if int(st.reg.PC()) >= len(st.mem) {
		return emu.MemoryAccessViolation
	}
	return emu.OK
}

func (st *state) Step() (diff *emu.Diff, c emu.Code) {
	if len(st.reg) < RegCount {
		return nil, emu.RegStateTooSmall
	}
	if c = st.Fetch(); c != emu.OK {
		return
	}
	fmt.Printf("lala\n")
	return
}

func (e *Emulator) Step(st *emu.State) (*emu.Diff, emu.Code) {
	st16 := &state{st.Mem, regState(st.Reg)}
	return st16.Step()
}
