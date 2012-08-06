package main

import (
	"fmt"

	"github.com/krasin/polyemu/arch/dcpu16"
	"github.com/krasin/polyemu/emu"
)

var zeroState = &emu.State{
	Mem: make([]byte, 2),
	Reg: make([]uint64, 30),
}

func findNops(e *dcpu16.Emulator, st *emu.State, in []uint16) (out []uint16) {
	for _, op := range in {
		st.Mem[0] = byte(op & 0xFF)
		st.Mem[1] = byte((op >> 8) & 0xFF)
		st.Reg[dcpu16.PC] = 0

		if diff, code := e.Step(st); code == emu.OK {
			if len(diff.Mem) == 0 && len(diff.Reg) == 1 && diff.Reg[dcpu16.PC] == 1 {
				out = append(out, op)
			}

		}
	}
	return
}

func main() {
	e := new(dcpu16.Emulator)
	fmt.Printf("Possible 2-byte nops (false positives are possible):\n")
	nops := make([]uint16, 65536)
	for i := 0; i < 65536; i++ {
		nops[i] = uint16(i)
	}
	states := []*emu.State{
		zeroState,
	}
	for _, st := range states {
		nops = findNops(e, st, nops)
	}
	for _, nop := range nops {
		mem := []byte{byte(nop & 0xFF), byte((nop >> 8) & 0xFF)}
		op, code := dcpu16.Disassemble(mem)
		fmt.Printf("0x%02x%02x %v", mem[1], mem[0], op)
		if code != emu.OK {
			fmt.Printf("(err. code=%d)", code)
		}
		fmt.Printf("\n")
	}
}
