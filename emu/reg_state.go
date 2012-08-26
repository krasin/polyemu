package emu

type Reg16State struct {
	a []uint64
	b []uint64
}

func NewReg16State(a []uint64) *Reg16State {
	res := &Reg16State{
		a: a,
		b: make([]uint64, len(a)),
	}
	copy(res.b, res.a)
	return res
}

func (r *Reg16State) Len() int {
	return len(r.a)
}

func (r *Reg16State) Get(ind int) uint16 {
	return uint16(r.b[ind])
}

func (r *Reg16State) Set(ind int, val uint16) {
	r.b[ind] = uint64(val)
}

func (r *Reg16State) Diff(diff []DiffPair) []DiffPair {
	for i, b := range r.b {
		if r.a[i] != b {
			diff = append(diff, DiffPair{uint64(i), b})
		}
	}
	return diff
}

func (r *Reg16State) Dec(ind int) uint16 {
	v := r.Get(ind)
	v--
	r.Set(ind, v)
	return v
}

func (r *Reg16State) Inc(ind int) uint16 {
	v := r.Get(ind)
	v++
	r.Set(ind, v)
	return v
}
