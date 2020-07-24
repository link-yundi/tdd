package ch02_document

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ExampleAdd() {
	sum := Add(2, 5)
	fmt.Println(sum)
	// Output: 7
}

func TestAdder(t *testing.T) {
	assertT := assert.New(t)
	got := Add(2, 5)
	want := 7
	assertT.Equal(want, got)
}



