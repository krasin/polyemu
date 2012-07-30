package dcpu16

import (
	"fmt"

	"github.com/krasin/polyemu/emu"
)

const (
	PC       = iota
	RegCount = iota
)

const (
	SPECIAL_OP = 0x00
	SET_OP     = 0x01
	ADD_OP     = 0x02
	SUB_OP     = 0x03
	MUL_OP     = 0x04
	MLI_OP     = 0x05
	DIV_OP     = 0x06
	DVI_OP     = 0x07
	MOD_OP     = 0x08
	MDI_OP     = 0x09
	AND_OP     = 0x0a
	BOR_OP     = 0x0b
	XOR_OP     = 0x0c
	SHR_OP     = 0x0d
	ASR_OP     = 0x0e
	SHL_OP     = 0x0f
	IFB_OP     = 0x10
	IFC_OP     = 0x11
	IFE_OP     = 0x12
	IFN_OP     = 0x13
	IFG_OP     = 0x14
	IFA_OP     = 0x15
	IFL_OP     = 0x16
	IFU_OP     = 0x17
	ADX_OP     = 0x1a
	SBX_OP     = 0x1b
	STI_OP     = 0x1e
	STD_OP     = 0x1f
)

type Emulator struct {
}

type memory []byte

func (m memory) At(ind uint16) uint16 {
	if int(2*ind) >= len(m) {
		return 0
	}
	if int(2*ind) == len(m)-1 {
		return uint16(m[2*ind])
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

	opcode int
}

func (st *state) fetchSpecial(v uint16) emu.Code {
	return emu.NotImplemented
}

func (st *state) Fetch() (c emu.Code) {
	v := st.mem.At(st.reg.PC())
	st.opcode = int(v & 0x1F)
	fmt.Printf("opcode: %x\n", st.opcode)
	switch st.opcode {
	case SPECIAL_OP:
		return st.fetchSpecial(v)
	case SET_OP:
	case ADD_OP:
	case SUB_OP:
	case MUL_OP:
	case MLI_OP:
	case DIV_OP:
	case DVI_OP:
	case MOD_OP:
	case MDI_OP:
	case AND_OP:
	case BOR_OP:
	case XOR_OP:
	case SHR_OP:
	case ASR_OP:
	case SHL_OP:
	case IFB_OP:
	case IFC_OP:
	case IFE_OP:
	case IFN_OP:
	case IFG_OP:
	case IFA_OP:
	case IFL_OP:
	case IFU_OP:
	case ADX_OP:
	case SBX_OP:
	case STI_OP:
	case STD_OP:
	default:
		return emu.InvalidOpcode
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
	st16 := &state{
		mem: memory(st.Mem),
		reg: regState(st.Reg),
	}
	return st16.Step()
}
