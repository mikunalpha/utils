package utils

func NewPtrString(s string) *string {
	return &s
}

type SliceOfString []string

func (sos SliceOfString) Unique() []string {
	mos := map[string]bool{}
	for _, s := range sos {
		mos[s] = false
	}
	uniqueSos := []string{}
	for _, s := range sos {
		if added, ok := mos[s]; ok && !added {
			uniqueSos = append(uniqueSos, s)
			mos[s] = true
		}
	}
	return uniqueSos
}

func (sos SliceOfString) InSlice(t string) bool {
	for _, s := range sos {
		if t == s {
			return true
		}
	}
	return false
}

type SliceOfPtrString []*string

func (sops SliceOfPtrString) Unique() []*string {
	mos := map[string]bool{}
	for _, ps := range sops {
		if ps == nil {
			continue
		}
		mos[*ps] = false
	}
	uniqueSos := []*string{}
	for _, ps := range sops {
		if added, ok := mos[*ps]; ok && !added {
			uniqueSos = append(uniqueSos, NewPtrString(*ps))
			mos[*ps] = true
		}
	}
	return uniqueSos
}

func (sops SliceOfPtrString) InSlice(t string) bool {
	for _, ps := range sops {
		if ps != nil && t == *ps {
			return true
		}
	}
	return false
}
