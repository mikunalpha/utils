package utils

func NewPtrUint(i uint) *uint {
	return &i
}

func NewPtrUint8(i uint8) *uint8 {
	return &i
}

func NewPtrUint16(i uint16) *uint16 {
	return &i
}

func NewPtrUint32(i uint32) *uint32 {
	return &i
}

func NewPtrUint64(i uint64) *uint64 {
	return &i
}

type SliceOfUint []uint

type SliceOfUint8 []uint8

type SliceOfUint16 []uint16

type SliceOfUint32 []uint32

type SliceOfUint64 []uint64

func (soui SliceOfUint) Unique() []uint {
	mos := map[uint]bool{}
	for _, ui := range soui {
		mos[ui] = false
	}
	uniqueSoui := []uint{}
	for _, ui := range soui {
		if added, ok := mos[ui]; ok && !added {
			uniqueSoui = append(uniqueSoui, ui)
			mos[ui] = true
		}
	}
	return uniqueSoui
}

func (soui SliceOfUint8) Unique() []uint8 {
	mos := map[uint8]bool{}
	for _, ui := range soui {
		mos[ui] = false
	}
	uniqueSoui := []uint8{}
	for _, ui := range soui {
		if added, ok := mos[ui]; ok && !added {
			uniqueSoui = append(uniqueSoui, ui)
			mos[ui] = true
		}
	}
	return uniqueSoui
}

func (soui SliceOfUint16) Unique() []uint16 {
	mos := map[uint16]bool{}
	for _, ui := range soui {
		mos[ui] = false
	}
	uniqueSoui := []uint16{}
	for _, ui := range soui {
		if added, ok := mos[ui]; ok && !added {
			uniqueSoui = append(uniqueSoui, ui)
			mos[ui] = true
		}
	}
	return uniqueSoui
}

func (soui SliceOfUint32) Unique() []uint32 {
	mos := map[uint32]bool{}
	for _, ui := range soui {
		mos[ui] = false
	}
	uniqueSoui := []uint32{}
	for _, ui := range soui {
		if added, ok := mos[ui]; ok && !added {
			uniqueSoui = append(uniqueSoui, ui)
			mos[ui] = true
		}
	}
	return uniqueSoui
}

func (soui SliceOfUint64) Unique() []uint64 {
	mos := map[uint64]bool{}
	for _, ui := range soui {
		mos[ui] = false
	}
	uniqueSoui := []uint64{}
	for _, ui := range soui {
		if added, ok := mos[ui]; ok && !added {
			uniqueSoui = append(uniqueSoui, ui)
			mos[ui] = true
		}
	}
	return uniqueSoui
}

func (soui SliceOfUint) InSlice(t uint) bool {
	for _, ui := range soui {
		if t == ui {
			return true
		}
	}
	return false
}

func (soui SliceOfUint8) InSlice(t uint8) bool {
	for _, ui := range soui {
		if t == ui {
			return true
		}
	}
	return false
}

func (soui SliceOfUint16) InSlice(t uint16) bool {
	for _, ui := range soui {
		if t == ui {
			return true
		}
	}
	return false
}

func (soui SliceOfUint32) InSlice(t uint32) bool {
	for _, ui := range soui {
		if t == ui {
			return true
		}
	}
	return false
}

func (soui SliceOfUint64) InSlice(t uint64) bool {
	for _, ui := range soui {
		if t == ui {
			return true
		}
	}
	return false
}

type SliceOfPtrUint []*uint

type SliceOfPtrUint8 []*uint8

type SliceOfPtrUint16 []*uint16

type SliceOfPtrUint32 []*uint32

type SliceOfPtrUint64 []*uint64

func (sopui SliceOfPtrUint) Unique() []*uint {
	moui := map[uint]bool{}
	for _, pui := range sopui {
		if pui == nil {
			continue
		}
		moui[*pui] = false
	}
	uniqueSopui := []*uint{}
	for _, pui := range sopui {
		if added, ok := moui[*pui]; ok && !added {
			uniqueSopui = append(uniqueSopui, NewPtrUint(*pui))
			moui[*pui] = true
		}
	}
	return uniqueSopui
}

func (sopui SliceOfPtrUint8) Unique() []*uint8 {
	moui := map[uint8]bool{}
	for _, pui := range sopui {
		if pui == nil {
			continue
		}
		moui[*pui] = false
	}
	uniqueSopui := []*uint8{}
	for _, pui := range sopui {
		if added, ok := moui[*pui]; ok && !added {
			uniqueSopui = append(uniqueSopui, NewPtrUint8(*pui))
			moui[*pui] = true
		}
	}
	return uniqueSopui
}

func (sopui SliceOfPtrUint16) Unique() []*uint16 {
	moui := map[uint16]bool{}
	for _, pui := range sopui {
		if pui == nil {
			continue
		}
		moui[*pui] = false
	}
	uniqueSopui := []*uint16{}
	for _, pui := range sopui {
		if added, ok := moui[*pui]; ok && !added {
			uniqueSopui = append(uniqueSopui, NewPtrUint16(*pui))
			moui[*pui] = true
		}
	}
	return uniqueSopui
}

func (sopui SliceOfPtrUint32) Unique() []*uint32 {
	moui := map[uint32]bool{}
	for _, pui := range sopui {
		if pui == nil {
			continue
		}
		moui[*pui] = false
	}
	uniqueSopui := []*uint32{}
	for _, pui := range sopui {
		if added, ok := moui[*pui]; ok && !added {
			uniqueSopui = append(uniqueSopui, NewPtrUint32(*pui))
			moui[*pui] = true
		}
	}
	return uniqueSopui
}

func (sopui SliceOfPtrUint64) Unique() []*uint64 {
	moui := map[uint64]bool{}
	for _, pui := range sopui {
		if pui == nil {
			continue
		}
		moui[*pui] = false
	}
	uniqueSopui := []*uint64{}
	for _, pui := range sopui {
		if added, ok := moui[*pui]; ok && !added {
			uniqueSopui = append(uniqueSopui, NewPtrUint64(*pui))
			moui[*pui] = true
		}
	}
	return uniqueSopui
}

func (sopui SliceOfPtrUint) InSlice(t uint) bool {
	for _, pui := range sopui {
		if pui != nil && t == *pui {
			return true
		}
	}
	return false
}

func (sopui SliceOfPtrUint8) InSlice(t uint8) bool {
	for _, pui := range sopui {
		if pui != nil && t == *pui {
			return true
		}
	}
	return false
}

func (sopui SliceOfPtrUint16) InSlice(t uint16) bool {
	for _, pui := range sopui {
		if pui != nil && t == *pui {
			return true
		}
	}
	return false
}

func (sopui SliceOfPtrUint32) InSlice(t uint32) bool {
	for _, pui := range sopui {
		if pui != nil && t == *pui {
			return true
		}
	}
	return false
}

func (sopui SliceOfPtrUint64) InSlice(t uint64) bool {
	for _, pui := range sopui {
		if pui != nil && t == *pui {
			return true
		}
	}
	return false
}
