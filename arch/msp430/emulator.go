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
	return uint16(m.Byte(2*int(ind))) + (uint16(m.Byte(2*int(ind)+1)) << 8)
}

func (m *memory) SetWord(ind uint16, val uint16) emu.Code {
	m.SetByte(2*int(ind), byte(val&0xFF))
	m.SetByte(2*int(ind)+1, byte((val>>8)&0xFF))
	return emu.OK
}

type regState struct {
	a []uint64
	b []uint64
}

func newRegState(a []uint64) *regState {
	res := &regState{
		a: a,
		b: make([]uint64, len(a)),
	}
	copy(res.b, res.a)
	return res
}

func (r *regState) Get(ind int) uint16 {
	return uint16(r.b[ind])
}

func (r *regState) Set(ind int, val uint16) {
	r.b[ind] = uint64(val)
}

func (r *regState) Diff(diff []emu.DiffPair) []emu.DiffPair {
	for i, b := range r.b {
		if r.a[i] != b {
			diff = append(diff, emu.DiffPair{uint64(i), b})
		}
	}
	return diff
}

func (r *regState) Dec(ind int) uint16 {
	v := r.Get(ind)
	v--
	r.Set(ind, v)
	return v
}

func (r *regState) Inc(ind int) uint16 {
	v := r.Get(ind)
	v++
	r.Set(ind, v)
	return v
}

type state struct {
	mem *memory
	reg *regState
}

func (st *state) doStep() (code emu.Code) {
	if len(st.reg.a) < RegCount {
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
		reg: newRegState(st.Reg),
	}
	return s.Step(diff)
}
