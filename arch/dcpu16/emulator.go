package dcpu16

import (
	//	"fmt"

	"github.com/krasin/polyemu/emu"
)

const (
	RA        = 0
	RB        = 1
	RC        = 2
	RX        = 3
	RY        = 4
	RZ        = 5
	RI        = 6
	RJ        = 7
	PC        = 8
	SP        = 9
	EX        = 10
	SKIP_FLAG = 11
	RegCount  = iota
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

	JSR_SP = 0x01
	INT_SP = 0x08
	IAG_SP = 0x09
	IAS_SP = 0x0a
	RFI_SP = 0x0b
	IAQ_SP = 0x0c
	HWN_SP = 0x10
	HWQ_SP = 0x11
	HWI_SP = 0x12

	REG_ARG           = 1 // register (A, B, C, ...) value
	REG_ADDR_ARG      = 2 // [register]
	REG_ADDR_WORD_ARG = 3 // [register + next word]
	PUSH_ARG          = 4 // (PUSH / [--SP])
	POP_ARG           = 5 // POP / [SP++]
	ADDR_WORD_ARG     = 6 // [next word]
	WORD_ARG          = 7 // next word (literal)
	LITERAL_ARG       = 8 // literal (-1..30)

	NOP_POST    = 0 // No post action
	INC_IJ_POST = 1 // Increment I and J
	DEC_IJ_POST = 2 // Decrement I and J
)

type addrMode int

type arg struct {
	mode addrMode
	val  uint16
	val2 uint16
}

type Emulator struct {
}

type memory struct {
	a    []byte
	diff []emu.DiffPair
}

func newMemory(a []byte) *memory {
	return &memory{
		a: a,
	}
}

func (m *memory) Diff(diff []emu.DiffPair) []emu.DiffPair {
	if len(m.diff) == 0 {
		return diff
	}
	for i, v := range m.diff {
		skip := false
		for j := i + 1; j < len(m.diff); j++ {
			if v.Ind == m.diff[j].Ind {
				skip = true
				break
			}
		}
		if !skip && m.a[v.Ind] != byte(v.Val) {
			diff = append(diff, v)
		}
	}
	return diff
}

func (m *memory) Byte(ind int) byte {
	//	fmt.Printf("memory.Byte(%d): ", ind)
	//	if v, ok := m.diff[uint64(ind)]; ok {
	//		fmt.Printf("%d\n", v)
	//		return v
	//	}
	if ind < len(m.a) {
		//		fmt.Printf("%d\n", m.a[ind])
		return m.a[ind]
	}
	//	fmt.Printf("0\n")
	return 0
}

func (m *memory) SetByte(ind int, val byte) {
	m.diff = append(m.diff, emu.DiffPair{uint64(ind), uint64(val)})
	//	fmt.Printf("memory.SetByte(%d, %d)\n", ind, val)
	//	cur := m.Byte(ind)
	//	if cur == val {
	//		return
	//	}
	//	if ind < len(m.a) && m.a[ind] == val {
	//		delete(m.diff, uint64(ind))
	//		return
	//	}
	//	m.diff[uint64(ind)] = val
}

func (m *memory) At(ind uint16) uint16 {
	//	fmt.Printf("memory.At(%d)\n", ind)
	return uint16(m.Byte(2*int(ind))) + (uint16(m.Byte(2*int(ind)+1)) << 8)
}

func (m *memory) Set(ind uint16, val uint16) emu.Code {
	//	fmt.Printf("memory.Set(%d, %d)\n", ind, val)
	m.SetByte(2*int(ind), byte(val&0xFF))
	m.SetByte(2*int(ind)+1, byte((val>>8)&0xFF))
	return emu.OK
}

type state struct {
	mem *memory
	reg *emu.Reg16State

	opcode int
	a      uint16
	b      uint16

	skipFetchB bool
	argA       arg
	argB       arg

	valA uint16
	valB uint16

	res        uint16
	postEffect int
	skipStore  bool
}

func (st *state) eatWord() uint16 {
	v := st.mem.At(st.reg.Get(PC))
	st.reg.Inc(PC)
	return v
}

func (st *state) pop() uint16 {
	v := st.mem.At(st.reg.Get(SP))
	st.reg.Inc(SP)
	return v
}

func (st *state) fetchFirst() (c emu.Code) {
	v := st.eatWord()
	st.opcode = int(v & 0x1F)
	//	fmt.Printf("opcode: 0x%x\n", st.opcode)
	switch st.opcode {
	case SPECIAL_OP:
		st.skipFetchB = true
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
	//	fmt.Printf("a: 0x%x, b: 0x%x\n", st.a, st.b)

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
	if !st.skipFetchB {
		if code = st.fetchB(); code != emu.OK {
			return
		}
	}
	//	fmt.Printf("st.argA: %+v, st.argB: %+v\n", st.argA, st.argB)
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
	//	fmt.Printf("st.valA = 0x%x\n", st.valA)

	if !st.skipFetchB {
		st.valB = st.loadVal(st.argB)
		//		fmt.Printf("st.valB = 0x%x\n", st.valB)
	}
	return emu.OK
}

func (st *state) execSpecial() (code emu.Code) {
	switch st.b {
	case JSR_SP:
		if code = st.push(st.reg.Get(PC)); code != emu.OK {
			return
		}
		st.reg.Set(PC, st.valA)
		st.skipStore = true
	case INT_SP, IAG_SP, IAS_SP, RFI_SP, IAQ_SP, HWN_SP, HWQ_SP, HWI_SP:
		return emu.Interrupt
	default:
		return emu.InvalidOpcode
	}
	return
}

func (st *state) exec() (code emu.Code) {
	switch st.opcode {
	case SPECIAL_OP:
		return st.execSpecial()
	case SET_OP:
		st.res = st.valA
	case ADD_OP:
		st.res = st.valB + st.valA
		st.reg.Set(EX, uint16((int(st.valB)+int(st.valA))>>16))
	case SUB_OP:
		st.res = st.valB - st.valA
		if st.valB < st.valA {
			st.reg.Set(EX, 0xFFFF)
		} else {
			st.reg.Set(EX, 0)
		}
	case MUL_OP:
		st.res = st.valB * st.valA
		st.reg.Set(EX, uint16(((uint64(st.valB)*uint64(st.valA))>>16)&0xFFFF))
	case MLI_OP:
		st.res = uint16(int16(st.valB) * int16(st.valA))
		st.reg.Set(EX, uint16(((int64(int16(st.valB))*int64(int16(st.valA)))>>16)&0xFFFF))
	case DIV_OP:
		if st.valA == 0 {
			st.res = 0
			st.reg.Set(EX, 0)
		} else {
			st.res = st.valB / st.valA
			st.reg.Set(EX, uint16(((uint64(st.valB)<<16)/uint64(st.valA))&0xFFFF))
		}
	case DVI_OP:
		if st.valA == 0 {
			st.res = 0
			st.reg.Set(EX, 0)
		} else {
			st.res = uint16(int16(st.valB) / int16(st.valA))
			st.reg.Set(EX, uint16(((int64(int16(st.valB))<<16)/int64(int16(st.valA)))&0xFFFF))
		}
	case MOD_OP:
		if st.valA == 0 {
			st.res = 0
		} else {
			st.res = st.valB % st.valA
		}
	case MDI_OP:
		if st.valA == 0 {
			st.res = 0
		} else {
			st.res = uint16(int16(st.valB) % int16(st.valA))
		}
	case AND_OP:
		st.res = st.valB & st.valA
	case BOR_OP:
		st.res = st.valB | st.valA
	case XOR_OP:
		st.res = st.valB ^ st.valA
	case SHR_OP:
		st.res = st.valB >> st.valA
		st.reg.Set(EX, uint16(((uint64(st.valB)<<16)>>st.valA)&0xFFFF))
	case ASR_OP:
		st.res = uint16(int64(int16(st.valB)) >> st.valA)
		st.reg.Set(EX, uint16(((int64(int16(st.valB))<<16)>>st.valA)&0xFFFF))
	case SHL_OP:
		st.res = st.valB << st.valA
		st.reg.Set(EX, uint16(((uint64(st.valB)<<st.valA)>>16)&0xFFFF))
	case IFB_OP:
		st.skipStore = true
		if !(st.valB&st.valA != 0) {
			// Ignore next instruction
			st.reg.Set(SKIP_FLAG, 1)
		}
	case IFC_OP:
		st.skipStore = true
		if !(st.valB&st.valA == 0) {
			// Ignore next instruction
			st.reg.Set(SKIP_FLAG, 1)
		}
	case IFE_OP:
		st.skipStore = true
		if !(st.valB == st.valA) {
			// Ignore next instruction
			st.reg.Set(SKIP_FLAG, 1)
		}
	case IFN_OP:
		st.skipStore = true
		if !(st.valB != st.valA) {
			// Ignore next instruction
			st.reg.Set(SKIP_FLAG, 1)
		}
	case IFG_OP:
		st.skipStore = true
		if !(st.valB > st.valA) {
			// Ignore next instruction
			st.reg.Set(SKIP_FLAG, 1)
		}
	case IFA_OP:
		st.skipStore = true
		if !(int16(st.valB) > int16(st.valA)) {
			// Ignore next instruction
			st.reg.Set(SKIP_FLAG, 1)
		}
	case IFL_OP:
		st.skipStore = true
		if !(st.valB < st.valA) {
			// Ignore next instruction
			st.reg.Set(SKIP_FLAG, 1)
		}
	case IFU_OP:
		st.skipStore = true
		if !(int16(st.valB) < int16(st.valA)) {
			// Ignore next instruction
			st.reg.Set(SKIP_FLAG, 1)
		}
	case ADX_OP:
		v := uint64(st.valB) + uint64(st.valA) + uint64(st.reg.Get(EX))
		st.res = uint16(v & 0xFFFF)
		if v > 0xFFFF {
			st.reg.Set(EX, 1)
		} else {
			st.reg.Set(EX, 0)
		}
	case SBX_OP:
		v := int64(st.valB) - int64(st.valA) + int64(st.reg.Get(EX))
		st.res = uint16(uint64(v) & 0xFFFF)
		if v < 0 {
			st.reg.Set(EX, 0xFFFF)
		} else {
			st.reg.Set(EX, 0)
		}
	case STI_OP:
		st.res = st.valA
		st.postEffect = INC_IJ_POST
	case STD_OP:
		st.res = st.valA
		st.postEffect = DEC_IJ_POST

	default:
		return emu.InvalidOpcode
	}
	//	fmt.Printf("st.res: 0x%x\n", st.res)
	return
}

func (st *state) push(val uint16) emu.Code {
	st.reg.Dec(SP)
	return st.mem.Set(st.reg.Get(SP), val)

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
		return st.push(st.res)
	case ADDR_WORD_ARG:
		return st.mem.Set(ar.val, st.res)
	case WORD_ARG, LITERAL_ARG:
		// do nothing
	default:
		return emu.NotImplemented
	}
	return emu.OK
}

func (st *state) store() emu.Code {
	return st.storeVal(st.argB)
}

func (st *state) post() emu.Code {
	switch st.postEffect {
	case NOP_POST:
	case INC_IJ_POST:
		st.reg.Inc(RI)
		st.reg.Inc(RJ)
	case DEC_IJ_POST:
		st.reg.Dec(RI)
		st.reg.Dec(RJ)
	default:
		panic("Not reachable")
	}
	return emu.OK
}

func (st *state) doStep() (code emu.Code) {
	if st.reg.Len() < RegCount {
		return emu.RegStateTooSmall
	}

	if code = st.fetch(); code != emu.OK {
		return
	}
	// Skip flag is set by conditional instructions, like, IFE or IFC,
	// in case if the condition was not satisfied
	if st.reg.Get(SKIP_FLAG) != 0 {
		st.reg.Set(SKIP_FLAG, 0)
		return
	}
	if code = st.load(); code != emu.OK {
		return
	}
	if code = st.exec(); code != emu.OK {
		return
	}
	if !st.skipStore {
		if code = st.store(); code != emu.OK {
			return
		}
	}
	if code = st.post(); code != emu.OK {
		return
	}
	return
}

func (st *state) Step(diff *emu.Diff) (code emu.Code) {
	if code = st.doStep(); code != emu.OK {
		return
	}
	diff.Mem = st.mem.Diff(diff.Mem)
	diff.Reg = st.reg.Diff(diff.Reg)

	return
}

func (e *Emulator) Step(st *emu.State, diff *emu.Diff) emu.Code {
	st16 := &state{
		mem: newMemory(st.Mem),
		reg: emu.NewReg16State(st.Reg),
	}
	return st16.Step(diff)
}
