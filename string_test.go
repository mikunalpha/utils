package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceOfString(t *testing.T) {
	sos := []string{"a", "b", "c"}
	dsos := []string{"a", "a", "b", "c"}
	// assert.Equal(t, "a", SliceOfString(sos).First(""), "they should be equal")
	// assert.Equal(t, "c", SliceOfString(sos).Last(""), "they should be equal")
	assert.Equal(t, true, SliceOfString(sos).InSlice("b"), "they should be equal")
	assert.Equal(t, sos, SliceOfString(dsos).Unique(), "they should be equal")
}

func TestSliceOfPtrString(t *testing.T) {
	sops := []*string{NewPtrString("a"), NewPtrString("b"), NewPtrString("c")}
	dsops := []*string{NewPtrString("a"), NewPtrString("a"), NewPtrString("b"), NewPtrString("c")}
	// assert.Equal(t, "a", *(SliceOfPtrString(sops).First("")), "they should be equal")
	// assert.Equal(t, "c", *(SliceOfPtrString(sops).Last("")), "they should be equal")
	assert.Equal(t, sops, SliceOfPtrString(dsops).Unique(), "they should be equal")
	assert.Equal(t, true, SliceOfPtrString(sops).InSlice("b"), "they should be equal")
}
