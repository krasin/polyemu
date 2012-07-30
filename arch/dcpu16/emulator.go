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

type memory []byte

func (m memory) At(ind uint16) uint16 {
	if int(2*ind) >= len(m) {
		return 0
	}
	return uint16(m[2*ind]) + (uint16(m[2*ind+1]) << 8)
}

type regState []uint64

func (r regState) PC() uint16 {
	return uint16(r[PC])
}

type state struct {
	mem memory
	reg regState
}

func (st *state) Fetch() (c emu.Code) {
	v := st.mem.At(st.reg.PC())
	opcode := v & 0x1F
	fmt.Printf("opcode: %x\n", opcode)

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
	st16 := &state{memory(st.Mem), regState(st.Reg)}
	return st16.Step()
}
