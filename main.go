package main

import (
	"fmt"

	"github.com/krasin/polyemu/arch/dcpu16"
	"github.com/krasin/polyemu/emu"
)

func main() {
	e := &dcpu16.Emulator{}
	st := &emu.State{
		Mem: make([]byte, 0), //128*1024),
		Reg: make([]uint64, 32),
	}
	if _, code := e.Step(st); code != emu.OK {
		fmt.Printf("code = %v\n", code)
	}
}
