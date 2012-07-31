package emu

import "fmt"

type State struct {
	Mem []byte
	Reg []uint64
}

type Diff struct {
	Mem []int32
	Reg []int32
}

type Emulator interface {
	Step(st *State) (*Diff, Code)
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
)
