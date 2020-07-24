package ch03_iteration

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeat(t *testing.T) {
	assert := assert.New(t)
	type input struct {
		character string
		repeatCount int
	}
	tests := []struct {
		name string
		input input
		want string
	}{
		// TODO: test cases
		{
			name: "0次",
			input: input{
				character: "a",
				repeatCount: 0,
			},
			want: "",
		},
		{
			name: "5次",
			input: input{
				character: "a",
				repeatCount: 5,
			},
			want: "aaaaa",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Repeat(test.input.character, test.input.repeatCount)
			want := test.want
			assert.Equal(want, got)
		})
	}
}


func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

