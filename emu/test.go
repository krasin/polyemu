package emu

type T interface {
	Errorf(fmt string, args ...interface{})
}

type Test struct {
	Mem     []byte
	Reg     []uint64
	WantReg []uint64
	N       int
}

func RunTest(t T, e Emulator, test Test) {
	st := &State{
		Mem: make([]byte, len(test.Mem)),
		Reg: make([]uint64, len(test.Reg)),
	}
	copy(st.Mem, test.Mem)
	copy(st.Reg, test.Reg)
	for i := 0; i < test.N; i++ {
		if _, code := e.Step(st); code != OK {
			t.Errorf("Execution failed at step %d with the following code: %v\n", i, code)
			return
		}
	}
	for i, got := range st.Reg {
		var want uint64
		if i < len(test.WantReg) {
			want = test.WantReg[i]
		}
		if got != want {
			t.Errorf("Reg[%d]: want 0x%x, got 0x%x\n", i, want, got)
		}
	}
}
