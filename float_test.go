package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceOfFloat(t *testing.T) {
	sofs := []float64{1, 2, 3}
	dsofs := []float64{1, 1, 2, 3}
	// assert.Equal(t, "a", SliceOfFloat(sofs).First(""), "they should be equal")
	// assert.Equal(t, "c", SliceOfFloat(sofs).Last(""), "they should be equal")
	assert.Equal(t, true, SliceOfFloat64(sofs).InSlice(2), "they should be equal")
	assert.Equal(t, sofs, SliceOfFloat64(dsofs).Unique(), "they should be equal")
}

func TestSliceOfPtrFloat(t *testing.T) {
	sopfs := []*float64{NewPtrFloat64(1), NewPtrFloat64(2), NewPtrFloat64(3)}
	dsopfs := []*float64{NewPtrFloat64(1), NewPtrFloat64(1), NewPtrFloat64(2), NewPtrFloat64(3)}
	// assert.Equal(t, "a", *(SliceOfPtrFloat(sopfs).First("")), "they should be equal")
	// assert.Equal(t, "c", *(SliceOfPtrFloat(sopfs).Last("")), "they should be equal")
	assert.Equal(t, sopfs, SliceOfPtrFloat64(dsopfs).Unique(), "they should be equal")
	assert.Equal(t, true, SliceOfPtrFloat64(sopfs).InSlice(2), "they should be equal")
}
