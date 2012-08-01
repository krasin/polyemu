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

	// SET Z, [Y]
	st.Mem[8] = 0xa1
	st.Mem[9] = 0x30

	// SET [Y+1], 5
	st.Mem[10] = 0x81
	st.Mem[11] = 0x9a
	st.Mem[12] = 0x01
	st.Mem[13] = 0x00

	// SET PUSH, 10
	st.Mem[14] = 0x01
	st.Mem[15] = 0x9b

	// SET [100], 12
	st.Mem[16] = 0xc1
	st.Mem[17] = 0xb7
	st.Mem[18] = 0x64
	st.Mem[19] = 0x00

	for i := 0; i < 8; i++ {
		fmt.Printf("Step %d\n", i)
		if _, code := e.Step(st); code != emu.OK {
			fmt.Printf("code = %v\n", code)
		}
		fmt.Printf("\n")
	}
}
