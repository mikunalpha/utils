package utils

func NewPtrBool(b bool) *bool {
	return &b
}

type SliceOfBool []bool

func (sob SliceOfBool) Unique() []bool {
	mob := map[bool]bool{}
	for _, b := range sob {
		mob[b] = false
	}
	uniqueSob := []bool{}
	for _, b := range sob {
		if added, ok := mob[b]; ok && !added {
			uniqueSob = append(uniqueSob, b)
			mob[b] = true
		}
	}
	return uniqueSob
}

func (sob SliceOfBool) InSlice(t bool) bool {
	for _, b := range sob {
		if t == b {
			return true
		}
	}
	return false
}

type SliceOfPtrBool []*bool

func (sopb SliceOfPtrBool) Unique() []*bool {
	mob := map[bool]bool{}
	for _, pb := range sopb {
		if pb == nil {
			continue
		}
		mob[*pb] = false
	}
	uniqueSob := []*bool{}
	for _, pb := range sopb {
		if added, ok := mob[*pb]; ok && !added {
			uniqueSob = append(uniqueSob, NewPtrBool(*pb))
			mob[*pb] = true
		}
	}
	return uniqueSob
}

func (sopb SliceOfPtrBool) InSlice(t bool) bool {
	for _, pb := range sopb {
		if pb != nil && t == *pb {
			return true
		}
	}
	return false
}
