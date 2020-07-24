package ch05

import (
	ch05 "github.com/link-yundi/tdd/ch05_struct_method_interface"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// 通过接口实现多态
func TestArea2(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name string
		shape ch05.Shape
		want float64
	}{
		// TODO: test cases
		{
			name: "rectangle area",
			shape: ch05.Rectangle{4, 7},
			want: float64(28),
		},
		{
			name: "circle area",
			shape: ch05.Circle{10},
			want: math.Pi * 100.0,
		},
		{
			name: "triangle",
			shape: ch05.Triangle{8, 5},
			want: 20.0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			assert.Equal(test.want, got)
		})
	}
}

