package main

import (
	"fmt"
	"math/rand"

	"github.com/krasin/polyemu/arch/dcpu16"
	"github.com/krasin/polyemu/emu"
)

var zeroState = &emu.State{
	Mem: make([]byte, 4),
	Reg: make([]uint64, 30),
}

func randRegState(seed int64) []uint64 {
	src := rand.NewSource(seed)
	rnd := rand.New(src)
	res := make([]uint64, 30)
	for i := range res {
		res[i] = uint64(rnd.Intn(65536))
	}
	res[dcpu16.SKIP_FLAG] = 0
	return res
}

func randMemState(mem []byte, seed int64) []byte {
	rnd := rand.New(rand.NewSource(seed))
	for i := range mem {
		mem[i] = byte(rnd.Intn(256))
	}
	return mem
}

func findNops(e *dcpu16.Emulator, st *emu.State, in []uint16) (out []uint16) {
	for _, op := range in {
		var diff *emu.Diff
		var code emu.Code
		st.Mem[0] = byte(op & 0xFF)
		st.Mem[1] = byte((op >> 8) & 0xFF)
		st.Mem[2] = st.Mem[0]
		st.Mem[3] = st.Mem[1]
		st.Reg[dcpu16.PC] = 0

		diff, code = e.Step(st)
		//			fmt.Printf("%+v\n", diff)
		if code != emu.OK || len(diff.Mem) != 0 || len(diff.Reg) != 1 || diff.Reg[dcpu16.PC] != 1 {
			continue
		}

		st.Apply(diff)
		diff, code = e.Step(st)
		//			fmt.Printf("%+v\n", diff)
		if code != emu.OK || len(diff.Mem) != 0 || len(diff.Reg) != 1 || diff.Reg[dcpu16.PC] != 2 {
			continue
		}

		out = append(out, op)
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

	nops = findNops(e, zeroState, nops)
	mem := make([]byte, 65536*2)
	for i := 0; i < 100; i++ {
		nops = findNops(e, &emu.State{
			Mem: randMemState(mem, int64(i)),
			Reg: randRegState(int64(i)),
		}, nops)

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
