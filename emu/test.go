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

func RunTest(testInd int, t T, e Emulator, test Test) {
	st := &State{
		Mem: make([]byte, len(test.Mem)),
		Reg: make([]uint64, len(test.Reg)),
	}
	copy(st.Mem, test.Mem)
	copy(st.Reg, test.Reg)
	for i := 0; i < test.N; i++ {
		diff := new(Diff)
		if code := e.Step(st, diff); code == OK {
			st.Apply(diff)
		} else {
			t.Errorf("test #%d: Execution failed at step %d with the following code: %v\n", testInd, i, code)
			return
		}
	}
	for i, got := range st.Reg {
		var want uint64
		if i < len(test.WantReg) {
			want = test.WantReg[i]
		}
		if got != want {
			t.Errorf("test #%d: Reg[%d]: want 0x%x, got 0x%x\n", testInd, i, want, got)
		}
	}
}
