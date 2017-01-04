package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceOfInt(t *testing.T) {
	soi := []int{1, 2, 3}
	dsoi := []int{1, 1, 2, 3}
	// assert.Equal(t, "a", SliceOfInt(soi).First(""), "they should be equal")
	// assert.Equal(t, "c", SliceOfInt(soi).Last(""), "they should be equal")
	assert.Equal(t, true, SliceOfInt(soi).InSlice(2), "they should be equal")
	assert.Equal(t, soi, SliceOfInt(dsoi).Unique(), "they should be equal")
}

func TestSliceOfPtrInt(t *testing.T) {
	sopis := []*int{NewPtrInt(1), NewPtrInt(2), NewPtrInt(3)}
	dsopis := []*int{NewPtrInt(1), NewPtrInt(1), NewPtrInt(2), NewPtrInt(3)}
	// assert.Equal(t, "a", *(SliceOfPtrInt(sopis).First("")), "they should be equal")
	// assert.Equal(t, "c", *(SliceOfPtrInt(sopis).Last("")), "they should be equal")
	assert.Equal(t, sopis, SliceOfPtrInt(dsopis).Unique(), "they should be equal")
	assert.Equal(t, true, SliceOfPtrInt(sopis).InSlice(2), "they should be equal")
}
