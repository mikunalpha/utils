package utils

func NewPtrInt(i int) *int {
	return &i
}

func NewPtrInt8(i int8) *int8 {
	return &i
}

func NewPtrInt16(i int16) *int16 {
	return &i
}

func NewPtrInt32(i int32) *int32 {
	return &i
}

func NewPtrInt64(i int64) *int64 {
	return &i
}

type SliceOfInt []int

type SliceOfInt8 []int8

type SliceOfInt16 []int16

type SliceOfInt32 []int32

type SliceOfInt64 []int64

func (soi SliceOfInt) Unique() []int {
	moi := map[int]bool{}
	for _, i := range soi {
		moi[i] = false
	}
	uniqueSoi := []int{}
	for _, i := range soi {
		if added, ok := moi[i]; ok && !added {
			uniqueSoi = append(uniqueSoi, i)
			moi[i] = true
		}
	}
	return uniqueSoi
}

func (soi SliceOfInt8) Unique() []int8 {
	moi := map[int8]bool{}
	for _, i := range soi {
		moi[i] = false
	}
	uniqueSoi := []int8{}
	for _, i := range soi {
		if added, ok := moi[i]; ok && !added {
			uniqueSoi = append(uniqueSoi, i)
			moi[i] = true
		}
	}
	return uniqueSoi
}

func (soi SliceOfInt16) Unique() []int16 {
	moi := map[int16]bool{}
	for _, i := range soi {
		moi[i] = false
	}
	uniqueSoi := []int16{}
	for _, i := range soi {
		if added, ok := moi[i]; ok && !added {
			uniqueSoi = append(uniqueSoi, i)
			moi[i] = true
		}
	}
	return uniqueSoi
}

func (soi SliceOfInt32) Unique() []int32 {
	moi := map[int32]bool{}
	for _, i := range soi {
		moi[i] = false
	}
	uniqueSoi := []int32{}
	for _, i := range soi {
		if added, ok := moi[i]; ok && !added {
			uniqueSoi = append(uniqueSoi, i)
			moi[i] = true
		}
	}
	return uniqueSoi
}

func (soi SliceOfInt64) Unique() []int64 {
	moi := map[int64]bool{}
	for _, i := range soi {
		moi[i] = false
	}
	uniqueSoi := []int64{}
	for _, i := range soi {
		if added, ok := moi[i]; ok && !added {
			uniqueSoi = append(uniqueSoi, i)
			moi[i] = true
		}
	}
	return uniqueSoi
}

func (soi SliceOfInt) InSlice(t int) bool {
	for _, i := range soi {
		if t == i {
			return true
		}
	}
	return false
}

func (soi SliceOfInt8) InSlice(t int8) bool {
	for _, i := range soi {
		if t == i {
			return true
		}
	}
	return false
}

func (soi SliceOfInt16) InSlice(t int16) bool {
	for _, i := range soi {
		if t == i {
			return true
		}
	}
	return false
}

func (soi SliceOfInt32) InSlice(t int32) bool {
	for _, i := range soi {
		if t == i {
			return true
		}
	}
	return false
}

func (soi SliceOfInt64) InSlice(t int64) bool {
	for _, i := range soi {
		if t == i {
			return true
		}
	}
	return false
}

type SliceOfPtrInt []*int

type SliceOfPtrInt8 []*int8

type SliceOfPtrInt16 []*int16

type SliceOfPtrInt32 []*int32

type SliceOfPtrInt64 []*int64

func (sopi SliceOfPtrInt) Unique() []*int {
	moi := map[int]bool{}
	for _, pi := range sopi {
		if pi == nil {
			continue
		}
		moi[*pi] = false
	}
	uniqueSopi := []*int{}
	for _, pi := range sopi {
		if added, ok := moi[*pi]; ok && !added {
			uniqueSopi = append(uniqueSopi, NewPtrInt(*pi))
			moi[*pi] = true
		}
	}
	return uniqueSopi
}

func (sopi SliceOfPtrInt8) Unique() []*int8 {
	moi := map[int8]bool{}
	for _, pi := range sopi {
		if pi == nil {
			continue
		}
		moi[*pi] = false
	}
	uniqueSopi := []*int8{}
	for _, pi := range sopi {
		if added, ok := moi[*pi]; ok && !added {
			uniqueSopi = append(uniqueSopi, NewPtrInt8(*pi))
			moi[*pi] = true
		}
	}
	return uniqueSopi
}

func (sopi SliceOfPtrInt16) Unique() []*int16 {
	moi := map[int16]bool{}
	for _, pi := range sopi {
		if pi == nil {
			continue
		}
		moi[*pi] = false
	}
	uniqueSopi := []*int16{}
	for _, pi := range sopi {
		if added, ok := moi[*pi]; ok && !added {
			uniqueSopi = append(uniqueSopi, NewPtrInt16(*pi))
			moi[*pi] = true
		}
	}
	return uniqueSopi
}

func (sopi SliceOfPtrInt32) Unique() []*int32 {
	moi := map[int32]bool{}
	for _, pi := range sopi {
		if pi == nil {
			continue
		}
		moi[*pi] = false
	}
	uniqueSopi := []*int32{}
	for _, pi := range sopi {
		if added, ok := moi[*pi]; ok && !added {
			uniqueSopi = append(uniqueSopi, NewPtrInt32(*pi))
			moi[*pi] = true
		}
	}
	return uniqueSopi
}

func (sopi SliceOfPtrInt64) Unique() []*int64 {
	moi := map[int64]bool{}
	for _, pi := range sopi {
		if pi == nil {
			continue
		}
		moi[*pi] = false
	}
	uniqueSopi := []*int64{}
	for _, pi := range sopi {
		if added, ok := moi[*pi]; ok && !added {
			uniqueSopi = append(uniqueSopi, NewPtrInt64(*pi))
			moi[*pi] = true
		}
	}
	return uniqueSopi
}

func (sopi SliceOfPtrInt) InSlice(t int) bool {
	for _, pi := range sopi {
		if pi != nil && t == *pi {
			return true
		}
	}
	return false
}

func (sopi SliceOfPtrInt8) InSlice(t int8) bool {
	for _, pi := range sopi {
		if pi != nil && t == *pi {
			return true
		}
	}
	return false
}

func (sopi SliceOfPtrInt16) InSlice(t int16) bool {
	for _, pi := range sopi {
		if pi != nil && t == *pi {
			return true
		}
	}
	return false
}

func (sopi SliceOfPtrInt32) InSlice(t int32) bool {
	for _, pi := range sopi {
		if pi != nil && t == *pi {
			return true
		}
	}
	return false
}

func (sopi SliceOfPtrInt64) InSlice(t int64) bool {
	for _, pi := range sopi {
		if pi != nil && t == *pi {
			return true
		}
	}
	return false
}
