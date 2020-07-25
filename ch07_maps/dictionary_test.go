package ch07_maps

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	v, ok := d[key]
	if !ok {
		return "", errors.New("could not find the word you are looking for")
	} else {
		return v, nil
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assert.Equal(t, want, got)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you are looking for"
		assert.Equal(t, want, err.Error())
	})
}

func TestError(t *testing.T) {
	err := errors.New("error test")
	check_str_1 := "error test"
	check_str_2 := "bad"

	fmt.Printf("%T\n", err.Error())
	// Output: string
	fmt.Println(err.Error())
	// Output: error test
	fmt.Println(err.Error() == check_str_1)
	// Output: true
	fmt.Println(err.Error() == check_str_2)
	// Output: false
}
