package test

import "testing"

func TestPerimeter(t *testing.T) {
	tests := []struct {
		name string
		width float64
		height float64
		perimeter float64
	}{
		// TODO: test cases
		{
			name: "2, 4",
			width: float64(2),
			height: float64(4),
			perimeter: float64(8),
		},
		{
			name: "4, 6",
			width: float64(4),
			height: float64(6),
			perimeter: float64(24),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}


