package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceOfUint(t *testing.T) {
	soui := []uint{1, 2, 3}
	dsoui := []uint{1, 1, 2, 3}
	// assert.Equal(t, "a", SliceOfUint(soui).First(""), "they should be equal")
	// assert.Equal(t, "c", SliceOfUint(soui).Last(""), "they should be equal")
	assert.Equal(t, true, SliceOfUint(soui).InSlice(2), "they should be equal")
	assert.Equal(t, soui, SliceOfUint(dsoui).Unique(), "they should be equal")
}

func TestSliceOfPtrUint(t *testing.T) {
	sopuis := []*uint{NewPtrUint(1), NewPtrUint(2), NewPtrUint(3)}
	dsopuis := []*uint{NewPtrUint(1), NewPtrUint(1), NewPtrUint(2), NewPtrUint(3)}
	// assert.Equal(t, "a", *(SliceOfPtrUint(sopuis).First("")), "they should be equal")
	// assert.Equal(t, "c", *(SliceOfPtrUint(sopuis).Last("")), "they should be equal")
	assert.Equal(t, sopuis, SliceOfPtrUint(dsopuis).Unique(), "they should be equal")
	assert.Equal(t, true, SliceOfPtrUint(sopuis).InSlice(2), "they should be equal")
}
