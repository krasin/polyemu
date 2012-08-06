package emu

import "fmt"

type State struct {
	Mem []byte
	Reg []uint64
}

func (st *State) Apply(diff *Diff) Code {
	for k, v := range diff.Mem {
		if k >= uint64(len(st.Mem)) {
			return MemoryAccessViolation
		}
		st.Mem[k] = v
	}
	for _, v := range diff.Reg {
		if v.Ind >= uint64(len(st.Reg)) {
			return RegStateTooSmall
		}
		st.Reg[v.Ind] = v.Val
	}
	return OK
}

type DiffPair struct {
	Ind uint64
	Val uint64
}

type Diff struct {
	Mem map[uint64]byte
	Reg []DiffPair
}

func (diff *Diff) Clear() {
	diff.Mem = nil
	diff.Reg = diff.Reg[:0]
}

func (diff *Diff) HasReg(ind, val uint64) bool {
	for _, v := range diff.Reg {
		if v.Ind == ind {
			return v.Val == val
		}
	}
	return false
}

type Emulator interface {
	Step(st *State, diff *Diff) Code
}

type Code uint32

func (c Code) String() string {
	switch c {
	case OK:
		return "OK"
	case RegStateTooSmall:
		return "RegStateTooSmall"
	case MemoryAccessViolation:
		return "MemoryAccessViolation"
	case InvalidOpcode:
		return "InvalidOpcode"
	case Interrupt:
		return "Interrupt"
	case NotImplemented:
		return "NotImplemented"
	case DecodeFailed:
		return "DecodeFailed"
	}
	return fmt.Sprintf("Code:%d", int(c))
}

const (
	OK                    = 0
	RegStateTooSmall      = 1
	MemoryAccessViolation = 2
	InvalidOpcode         = 3
	Interrupt             = 4
	NotImplemented        = 5
	DecodeFailed          = 6
)
