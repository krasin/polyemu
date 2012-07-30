package emu

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

const (
	OK                    = 0
	RegStateTooSmall      = 1
	MemoryAccessViolation = 2
	InvalidOpcode         = 3
	Interrupt             = 4
)
