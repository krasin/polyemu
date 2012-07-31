package main

import (
	"fmt"

	"github.com/krasin/polyemu/arch/dcpu16"
	"github.com/krasin/polyemu/emu"
)

func main() {
	e := &dcpu16.Emulator{}
	st := &emu.State{
		Mem: make([]byte, 128*1024),
		Reg: make([]uint64, 32),
	}
	// SET X, 1
	st.Mem[0] = 0x61
	st.Mem[1] = 0x88
	if _, code := e.Step(st); code != emu.OK {
		fmt.Printf("code = %v\n", code)
	}
}
