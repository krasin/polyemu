package dcpu16

import (
	"fmt"

	"github.com/krasin/polyemu/emu"
)

const (
	RA       = iota
	RB       = iota
	RC       = iota
	RX       = iota
	RY       = iota
	RZ       = iota
	RI       = iota
	RJ       = iota
	PC       = iota
	SP       = iota
	EX       = iota
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

	REG_ARG           = 1 // register (A, B, C, ...) value
	REG_ADDR_ARG      = 2 // [register]
	REG_ADDR_WORD_ARG = 3 // [register + next word]
	PUSH_ARG          = 4 // (PUSH / [--SP])
	POP_ARG           = 5 // POP / [SP++]
	ADDR_WORD_ARG     = 6 // [next word]
	WORD_ARG          = 7 // next word (literal)
	LITERAL_ARG       = 8 // literal (-1..30)
)

type addrMode int

type arg struct {
	mode addrMode
	val  int
	val2 int
}

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

func (r regState) SetPC(val uint16) {
	r[PC] = uint64(val)
}

type state struct {
	mem memory
	reg regState

	opcode int
	a      int
	b      int

	argA arg
	argB arg
}

func (st *state) fetchSpecial(v int) emu.Code {
	return emu.NotImplemented
}

func (st *state) eatWord() int {
	v := int(st.mem.At(st.reg.PC()))
	st.reg.SetPC(st.reg.PC() + 1)
	return v
}

func (st *state) fetchFirst() (c emu.Code) {
	v := st.eatWord()
	st.opcode = v & 0x1F
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
	st.a = (v >> 10) & 0x3F
	st.b = (v >> 5) & 0x1F
	fmt.Printf("a: %x, b: %x\n", st.a, st.b)

	return emu.OK
}

func (st *state) fetchCommonArg(v int) (ar arg, ok bool) {
	switch v {
	// register
	case 0x00:
		return arg{REG_ARG, RA, 0}, true
	case 0x01:
		return arg{REG_ARG, RB, 0}, true
	case 0x02:
		return arg{REG_ARG, RC, 0}, true
	case 0x03:
		return arg{REG_ARG, RX, 0}, true
	case 0x04:
		return arg{REG_ARG, RY, 0}, true
	case 0x05:
		return arg{REG_ARG, RZ, 0}, true
	case 0x06:
		return arg{REG_ARG, RI, 0}, true
	case 0x07:
		return arg{REG_ARG, RJ, 0}, true

	// [register]
	case 0x08:
		return arg{REG_ADDR_ARG, RA, 0}, true
	case 0x09:
		return arg{REG_ADDR_ARG, RB, 0}, true
	case 0x0A:
		return arg{REG_ADDR_ARG, RC, 0}, true
	case 0x0B:
		return arg{REG_ADDR_ARG, RX, 0}, true
	case 0x0C:
		return arg{REG_ADDR_ARG, RY, 0}, true
	case 0x0D:
		return arg{REG_ADDR_ARG, RZ, 0}, true
	case 0x0E:
		return arg{REG_ADDR_ARG, RI, 0}, true
	case 0x0F:
		return arg{REG_ADDR_ARG, RJ, 0}, true

	// [register+word]
	case 0x10:
		return arg{REG_ADDR_WORD_ARG, RA, st.eatWord()}, true
	case 0x11:
		return arg{REG_ADDR_WORD_ARG, RB, st.eatWord()}, true
	case 0x12:
		return arg{REG_ADDR_WORD_ARG, RC, st.eatWord()}, true
	case 0x13:
		return arg{REG_ADDR_WORD_ARG, RX, st.eatWord()}, true
	case 0x14:
		return arg{REG_ADDR_WORD_ARG, RY, st.eatWord()}, true
	case 0x15:
		return arg{REG_ADDR_WORD_ARG, RZ, st.eatWord()}, true
	case 0x16:
		return arg{REG_ADDR_WORD_ARG, RI, st.eatWord()}, true
	case 0x17:
		return arg{REG_ADDR_WORD_ARG, RJ, st.eatWord()}, true

	// Other
	case 0x19:
		return arg{REG_ADDR_ARG, SP, 0}, true
	case 0x1A:
		return arg{REG_ADDR_WORD_ARG, SP, st.eatWord()}, true
	case 0x1B:
		return arg{REG_ARG, SP, 0}, true
	case 0x1C:
		return arg{REG_ARG, PC, 0}, true
	case 0x1D:
		return arg{REG_ARG, EX, 0}, true
	case 0x1E:
		return arg{ADDR_WORD_ARG, st.eatWord(), 0}, true
	case 0x1F:
		return arg{WORD_ARG, st.eatWord(), 0}, true
	}

	// Not handled
	return
}

func (st *state) fetchA() (code emu.Code) {
	if arg, ok := st.fetchCommonArg(st.a); ok {
		st.argA = arg
		return
	}
	if st.a == 0x18 {
		st.argA = arg{POP_ARG, 0, 0}
		return
	}
	if st.a >= 0x20 && st.a <= 0x3F {
		st.argA = arg{LITERAL_ARG, st.a - 0x20 - 1, 0}
		return
	}
	return emu.DecodeFailed
}

func (st *state) fetchB() (code emu.Code) {
	if arg, ok := st.fetchCommonArg(st.b); ok {
		st.argB = arg
		return
	}
	if st.b == 0x18 {
		st.argB = arg{PUSH_ARG, 0, 0}
		return
	}
	return emu.DecodeFailed
}

func (st *state) fetch() (code emu.Code) {
	if code = st.fetchFirst(); code != emu.OK {
		return
	}
	if code = st.fetchA(); code != emu.OK {
		return
	}
	if code = st.fetchB(); code != emu.OK {
		return
	}
	return emu.OK
}

func (st *state) Step() (diff *emu.Diff, c emu.Code) {
	if len(st.reg) < RegCount {
		return nil, emu.RegStateTooSmall
	}
	if c = st.fetch(); c != emu.OK {
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
