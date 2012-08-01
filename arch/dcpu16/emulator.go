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
	val  uint16
	val2 uint16
}

type Emulator struct {
}

type memory []byte

func (m memory) At(ind uint16) uint16 {
	fmt.Printf("memory.At(%d)\n", ind)
	if 2*int(ind) >= len(m) {
		return 0
	}
	if 2*int(ind) == len(m)-1 {
		return uint16(m[2*int(ind)])
	}
	return uint16(m[2*int(ind)]) + (uint16(m[2*int(ind)+1]) << 8)
}

func (m memory) Set(ind uint16, val uint16) emu.Code {
	fmt.Printf("memory.Set(%d, %d)\n", ind, val)
	if 2*int(ind)+1 >= len(m) {
		return emu.MemoryAccessViolation
	}
	m[2*int(ind)] = byte(val & 0xFF)
	m[2*int(ind)+1] = byte((val >> 8) & 0xFF)
	return emu.OK
}

type regState []uint64

func (r regState) Get(ind int) uint16 {
	return uint16(r[ind])
}

func (r regState) Set(ind int, val uint16) {
	r[ind] = uint64(val)
}

func (r regState) Dec(ind int) uint16 {
	v := uint16(r[ind])
	v--
	r[ind] = uint64(v)
	return v
}

func (r regState) PC() uint16 {
	return uint16(r[PC])
}

func (r regState) SetPC(val uint16) {
	r[PC] = uint64(val)
}

func (r regState) SP() uint16 {
	return uint16(r[SP])
}

func (r regState) SetSP(val uint16) {
	r[SP] = uint64(val)
}

func (r regState) EX() uint16 {
	return uint16(r[EX])
}

func (r regState) SetEX(val uint16) {
	r[EX] = uint64(val)
}

type state struct {
	mem memory
	reg regState

	opcode int
	a      uint16
	b      uint16

	argA arg
	argB arg

	valA uint16
	valB uint16

	res uint16
}

func (st *state) fetchSpecial(v uint16) emu.Code {
	return emu.NotImplemented
}

func (st *state) eatWord() uint16 {
	v := st.mem.At(st.reg.PC())
	st.reg.SetPC(st.reg.PC() + 1)
	return v
}

func (st *state) pop() uint16 {
	v := st.mem.At(st.reg.SP())
	st.reg.SetSP(st.reg.SP() + 1)
	return v
}

func (st *state) fetchFirst() (c emu.Code) {
	v := st.eatWord()
	st.opcode = int(v & 0x1F)
	fmt.Printf("opcode: 0x%x\n", st.opcode)
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
	fmt.Printf("a: 0x%x, b: 0x%x\n", st.a, st.b)

	return emu.OK
}

func (st *state) fetchCommonArg(v uint16) (ar arg, ok bool) {
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
	fmt.Printf("st.argA: %+v, st.argB: %+v\n", st.argA, st.argB)
	return emu.OK
}

func (st *state) loadVal(ar arg) uint16 {
	switch ar.mode {
	case REG_ARG:
		return st.reg.Get(int(ar.val))
	case REG_ADDR_ARG:
		return st.mem.At(st.reg.Get(int(ar.val)))
	case REG_ADDR_WORD_ARG:
		return st.mem.At(st.reg.Get(int(ar.val)) + ar.val2)
	case POP_ARG:
		return st.pop()
	case PUSH_ARG:
		// Nothing to load
		return 0
	case ADDR_WORD_ARG:
		return st.mem.At(ar.val)
	case WORD_ARG:
		return ar.val
	case LITERAL_ARG:
		return ar.val
	}
	panic("not reachable")
}

func (st *state) load() (code emu.Code) {
	st.valA = st.loadVal(st.argA)
	fmt.Printf("st.valA = 0x%x\n", st.valA)
	st.valB = st.loadVal(st.argB)
	fmt.Printf("st.valB = 0x%x\n", st.valB)
	return emu.OK
}

func (st *state) exec() (code emu.Code) {
	switch st.opcode {
	case SET_OP:
		st.res = st.valA
	case ADD_OP:
		st.res = st.valB + st.valA
		st.reg.SetEX(uint16((int(st.valB) + int(st.valA)) >> 16))
	case SUB_OP:
		st.res = st.valB - st.valA
		if st.valB < st.valA {
			st.reg.SetEX(0xFFFF)
		} else {
			st.reg.SetEX(0)
		}
	case MUL_OP:
		st.res = st.valB * st.valA
		st.reg.SetEX(uint16(((uint64(st.valB) * uint64(st.valA)) >> 16) & 0xFFFF))
	default:
		return emu.NotImplemented
	}
	fmt.Printf("st.res: %x\n", st.res)
	return
}

func (st *state) storeVal(ar arg) emu.Code {
	switch ar.mode {
	case REG_ARG:
		st.reg.Set(int(ar.val), st.res)
	case REG_ADDR_ARG:
		return st.mem.Set(st.reg.Get(int(ar.val)), st.res)
	case REG_ADDR_WORD_ARG:
		return st.mem.Set(st.reg.Get(int(ar.val))+ar.val2, st.res)
	case POP_ARG:
		panic("not reachable")
	case PUSH_ARG:
		st.reg.Dec(SP)
		return st.mem.Set(st.reg.Get(SP), st.res)
	default:
		return emu.NotImplemented
	}
	return emu.OK
}

func (st *state) store() emu.Code {
	return st.storeVal(st.argB)
}

func (st *state) Step() (diff *emu.Diff, c emu.Code) {
	if len(st.reg) < RegCount {
		return nil, emu.RegStateTooSmall
	}
	if c = st.fetch(); c != emu.OK {
		return
	}
	if c = st.load(); c != emu.OK {
		return
	}
	if c = st.exec(); c != emu.OK {
		return
	}
	if c = st.store(); c != emu.OK {
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
