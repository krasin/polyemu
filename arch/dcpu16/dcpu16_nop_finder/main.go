package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"

	"github.com/krasin/polyemu/arch/dcpu16"
	"github.com/krasin/polyemu/emu"
)

var n = flag.Int("n", 100, "Number of random tries. 100k works about 20 minutes")
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

var zeroState = &emu.State{
	Mem: make([]byte, 1<<17),
	Reg: make([]uint64, 30),
}

func randRegState(seed int64) []uint64 {
	src := rand.NewSource(seed)
	rnd := rand.New(src)
	res := make([]uint64, 30)
	var v uint32
	var k int
	for i := range res {
		if k == 0 {
			v = rnd.Uint32()
			k = 2
		}
		res[i] = uint64(v & 0xFFFF)
		v >>= 16
		k--
	}
	res[dcpu16.SKIP_FLAG] = 0
	return res
}

func randMemState(mem []byte, seed int64) []byte {
	rnd := rand.New(rand.NewSource(seed))
	var v uint64
	var k int
	for i := range mem {
		if k == 0 {
			v = uint64(rnd.Int63n(0xFFFFFFFFFFFFFF))
			k = 7
		}
		mem[i] = byte(v & 0xFF)
		v >>= 8
		k--
	}
	return mem
}

func findNops(e *dcpu16.Emulator, st *emu.State, pc uint16, in []uint16) (out []uint16) {
	diff := new(emu.Diff)
	for _, op := range in {

		var code emu.Code
		st.Mem[2*int(pc)] = byte(op & 0xFF)
		st.Mem[2*int(pc)+1] = byte((op >> 8) & 0xFF)
		st.Mem[2*int(pc+1)] = st.Mem[2*int(pc)]
		st.Mem[2*int(pc+1)+1] = st.Mem[2*int(pc)+1]
		st.Reg[dcpu16.PC] = uint64(pc)

		diff.Clear()
		code = e.Step(st, diff)
		//			fmt.Printf("%+v\n", diff)

		if code != emu.OK || len(diff.Mem) != 0 || len(diff.Reg) != 1 || !diff.HasReg(dcpu16.PC, uint64(pc+1)) {
			continue
		}

		st.Apply(diff)
		diff.Clear()
		code = e.Step(st, diff)
		//			fmt.Printf("%+v\n", diff)
		if code != emu.OK || len(diff.Mem) != 0 || len(diff.Reg) != 1 || !diff.HasReg(dcpu16.PC, uint64(pc+2)) {
			continue
		}

		out = append(out, op)
	}
	return
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	e := new(dcpu16.Emulator)
	fmt.Printf("Possible 2-byte nops (false positives are possible):\n")
	nops := make([]uint16, 65536)
	for i := 0; i < 65536; i++ {
		nops[i] = uint16(i)
	}

	nops = findNops(e, zeroState, 0, nops)
	nops = findNops(e, zeroState, 0xFFFF, nops)
	nops = findNops(e, zeroState, 0x1234, nops)
	mem := make([]byte, 65536*2)
	for i := 0; i < *n; i++ {
		nops = findNops(e, &emu.State{
			Mem: randMemState(mem, int64(i)),
			Reg: randRegState(int64(i)),
		}, uint16(i), nops)
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
