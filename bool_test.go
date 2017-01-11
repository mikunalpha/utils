package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceOfBool(t *testing.T) {
	sob := []bool{true, false}
	dsob := []bool{true, true, false}
	// assert.Equal(t, "a", SliceOfBool(sob).First(""), "they should be equal")
	// assert.Equal(t, "c", SliceOfBool(sob).Last(""), "they should be equal")
	assert.Equal(t, true, SliceOfBool(sob).InSlice(false), "they should be equal")
	assert.Equal(t, sob, SliceOfBool(dsob).Unique(), "they should be equal")
}

func TestSliceOfPtrBool(t *testing.T) {
	sopb := []*bool{NewPtrBool(true), NewPtrBool(false)}
	dsopb := []*bool{NewPtrBool(true), NewPtrBool(true), NewPtrBool(false)}
	// assert.Equal(t, "a", *(SliceOfPtrBool(sopb).First("")), "they should be equal")
	// assert.Equal(t, "c", *(SliceOfPtrBool(sopb).Last("")), "they should be equal")
	assert.Equal(t, true, SliceOfPtrBool(sopb).InSlice(true), "they should be equal")
	assert.Equal(t, sopb, SliceOfPtrBool(dsopb).Unique(), "they should be equal")
}
