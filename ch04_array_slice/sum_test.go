package ch04_array_slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name string
		input []int
		want int
	}{
		// TODO: test cases
		{
			name: "sum of [1, 2, 3, 4, 5]",
			input: []int{1, 2, 3, 4, 5},
			want: 15,
		},
		{
			name: "sum of [10,]",
			input: []int{10},
			want: 10,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Sum(test.input)
			want := test.want
			assert.Equal(want, got)
		})
	}
}

dasdasdasdasdsad
