package main

import (
	"github.com/krasin/polyemu/arch/dcpu16"
	"github.com/krasin/polyemu/emu"
)

func main() {
	e := new(dcpu16.Emulator)
	e.Step(new(emu.State))
}
