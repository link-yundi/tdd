package ch05

import (
	. "github.com/link-yundi/tdd/ch05_struct_method_interface"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerimeter(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name string
		//width float64
		//height float64
		rectangle Rectangle
		want      float64
	}{
		// TODO: test cases
		{
			name: "2, 4",
			//width: float64(2),
			//height: float64(4),
			rectangle: Rectangle{2, 4},
			want:      float64(12),
		},
		{
			name: "4, 6",
			//width: float64(4),
			//height: float64(6),
			rectangle: Rectangle{4, 6},
			want:      float64(20),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Perimeter(test.rectangle)
			assert.Equal(test.want, got)
		})
	}
}
