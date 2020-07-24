package ch05

import (
	. "github.com/link-yundi/tdd/ch05_struct_method_interface"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArea(t *testing.T) {
	got := Area(Rectangle{2, 4})
	want := float64(8)
	assert.Equal(t, want, got)
}
