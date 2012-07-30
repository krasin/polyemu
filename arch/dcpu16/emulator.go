package dcpu16

import (
	"fmt"

	"github.com/krasin/polyemu/emu"
)

type Emulator struct {
}

type state emu.State

func (st *state) Step() (diff emu.Diff, status emu.Status) {
	fmt.Printf("lala\n")
	return
}

func (e *Emulator) Step(st *emu.State) (emu.Diff, emu.Status) {
	return ((*state)(st)).Step()
}
