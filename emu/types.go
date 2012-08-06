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
	for k, v := range diff.Reg {
		if k >= uint64(len(st.Reg)) {
			return RegStateTooSmall
		}
		st.Reg[k] = v
	}
	return OK
}

type Diff struct {
	Mem map[uint64]byte
	Reg map[uint64]uint64
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
