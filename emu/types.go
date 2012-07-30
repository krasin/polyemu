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
	Step(st *State) (Diff, Status)
}

type Status uint32

const (
	OK                      = 0
	RegStateTooSmall        = 0
	MemoryAccessViolation   = 2
	InstructionDecodeFailed = 3
	Interrupt               = 4
)
