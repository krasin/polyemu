package main

import (
	"fmt"

	"github.com/krasin/polyemu/arch/dcpu16"
	"github.com/krasin/polyemu/emu"
)

func main() {
	st := &emu.State{
		Mem: make([]byte, 2),
		Reg: make([]uint64, 30),
	}
	e := new(dcpu16.Emulator)
	fmt.Printf("Possible 2-byte nops (false positives are possible):\n")
	for i := 0; i < 256; i++ {
		st.Mem[0] = byte(i)
		for j := 0; j < 256; j++ {
			st.Mem[1] = byte(j)
			if diff, code := e.Step(st); code == emu.OK {
				if len(diff.Mem) == 0 && len(diff.Reg) == 1 && diff.Reg[dcpu16.PC] == 1 {
					op, code := dcpu16.Disassemble(st.Mem[:2])

					fmt.Printf("0x%02x%02x %v", j, i, op)
					if code != emu.OK {
						fmt.Printf("(err. code=%d)", code)
					}
					fmt.Printf("\n")

					//					fmt.Printf("DAT 0x%02x%02x\n", j, i)
				}
			}
		}
	}
}
