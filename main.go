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

	// MUL X, X
	st.Mem[2] = 0x64
	st.Mem[3] = 0x0c

	// SET Y, 30
	st.Mem[4] = 0x81
	st.Mem[5] = 0xfc

	// SET [Y], 4
	st.Mem[6] = 0x81
	st.Mem[7] = 0x95

	for i := 0; i < 4; i++ {
		fmt.Printf("Step %d\n", i)
		if _, code := e.Step(st); code != emu.OK {
			fmt.Printf("code = %v\n", code)
		}
	}
}
