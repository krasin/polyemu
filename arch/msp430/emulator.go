package msp430

import (
	"github.com/krasin/polyemu/emu"
)

const (
	RegCount = 16
)

type memory struct {
	a    []byte
	diff []emu.DiffPair
}

func newMemory(a []byte) *memory {
	return &memory{
		a: a,
	}
}

func (m *memory) Diff(diff []emu.DiffPair) []emu.DiffPair {
	if len(m.diff) == 0 {
		return diff
	}
	for i, v := range m.diff {
		skip := false
		for j := i + 1; j < len(m.diff); j++ {
			if v.Ind == m.diff[j].Ind {
				skip = true
				break
			}
		}
		if !skip && m.a[v.Ind] != byte(v.Val) {
			diff = append(diff, v)
		}
	}
	return diff
}

func (m *memory) Byte(ind int) byte {
	if ind < len(m.a) {
		return m.a[ind]
	}
	return 0
}

func (m *memory) SetByte(ind int, val byte) {
	m.diff = append(m.diff, emu.DiffPair{uint64(ind), uint64(val)})
}

func (m *memory) Word(ind uint16) uint16 {
	return uint16(m.Byte(int(2*ind))) + (uint16(m.Byte(int(2*ind+1))) << 8)
}

func (m *memory) SetWord(ind uint16, val uint16) emu.Code {
	m.SetByte(int(2*ind), byte(val&0xFF))
	m.SetByte(int(2*ind+1), byte((val>>8)&0xFF))
	return emu.OK
}

type state struct {
	mem *memory
	reg *emu.Reg16State
}

func (st *state) doStep() (code emu.Code) {
	if st.reg.Len() < RegCount {
		return emu.RegStateTooSmall
	}
	return emu.NotImplemented
}

func (st *state) Step(diff *emu.Diff) (code emu.Code) {
	if code = st.doStep(); code != emu.OK {
		return
	}
	diff.Mem = st.mem.Diff(diff.Mem)
	diff.Reg = st.reg.Diff(diff.Reg)

	return
}

type Emulator struct {
}

func (e *Emulator) Step(st *emu.State, diff *emu.Diff) emu.Code {
	s := &state{
		mem: newMemory(st.Mem),
		reg: emu.NewReg16State(st.Reg),
	}
	return s.Step(diff)
}
