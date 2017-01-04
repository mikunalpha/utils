package utils

func NewPtrFloat32(f float32) *float32 {
	return &f
}

func NewPtrFloat64(f float64) *float64 {
	return &f
}

type SliceOfFloat32 []float32

type SliceOfFloat64 []float64

func (sof SliceOfFloat32) Unique() []float32 {
	mof := map[float32]bool{}
	for _, f := range sof {
		mof[f] = false
	}
	uniqueSof := []float32{}
	for _, f := range sof {
		if added, ok := mof[f]; ok && !added {
			uniqueSof = append(uniqueSof, f)
			mof[f] = true
		}
	}
	return uniqueSof
}

func (sof SliceOfFloat64) Unique() []float64 {
	mof := map[float64]bool{}
	for _, f := range sof {
		mof[f] = false
	}
	uniqueSof := []float64{}
	for _, f := range sof {
		if added, ok := mof[f]; ok && !added {
			uniqueSof = append(uniqueSof, f)
			mof[f] = true
		}
	}
	return uniqueSof
}

func (sof SliceOfFloat32) InSlice(t float32) bool {
	for _, f := range sof {
		if t == f {
			return true
		}
	}
	return false
}

func (sof SliceOfFloat64) InSlice(t float64) bool {
	for _, f := range sof {
		if t == f {
			return true
		}
	}
	return false
}

type SliceOfPtrFloat32 []*float32

type SliceOfPtrFloat64 []*float64

func (sopf SliceOfPtrFloat32) Unique() []*float32 {
	mof := map[float32]bool{}
	for _, pf := range sopf {
		if pf == nil {
			continue
		}
		mof[*pf] = false
	}
	uniqueSopf := []*float32{}
	for _, pf := range sopf {
		if added, ok := mof[*pf]; ok && !added {
			uniqueSopf = append(uniqueSopf, NewPtrFloat32(*pf))
			mof[*pf] = true
		}
	}
	return uniqueSopf
}

func (sopf SliceOfPtrFloat64) Unique() []*float64 {
	mof := map[float64]bool{}
	for _, pf := range sopf {
		if pf == nil {
			continue
		}
		mof[*pf] = false
	}
	uniqueSopf := []*float64{}
	for _, pf := range sopf {
		if added, ok := mof[*pf]; ok && !added {
			uniqueSopf = append(uniqueSopf, NewPtrFloat64(*pf))
			mof[*pf] = true
		}
	}
	return uniqueSopf
}

func (sopf SliceOfPtrFloat32) InSlice(t float32) bool {
	for _, pf := range sopf {
		if pf != nil && t == *pf {
			return true
		}
	}
	return false
}

func (sopf SliceOfPtrFloat64) InSlice(t float64) bool {
	for _, pf := range sopf {
		if pf != nil && t == *pf {
			return true
		}
	}
	return false
}
