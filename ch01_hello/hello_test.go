package ch01_hello


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	//english = "English"
	french = "French"
	spanish = "Spanish"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloProfix = "Bonjour, "
)

func Hello(name, language string) string {
	if name == ""{
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloProfix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
//func TestHello(t *testing.T) {
//got := Hello("ZhangYundi")
//want := "Hello ZhangYundi"
//t.Logf("%[1]q, %[1]s", got)
//assert.Equal(t, want, got)
//	t.Run()

//}


func TestHello(t *testing.T) {
	// define assert func
	//assertCorrectMessage := func(t *testing.T, got, want string) {
	//	t.Helper()
	//	if got != want {
	//		t.Errorf("got '%q', want '%q'", got, want)
	//	}
	//}
	assertT := assert.New(t)
	type input struct {
		name string
		language string
	}
	tests := []struct {
		name string
		input input
		want string
	}{
		// TODO: test cases
		{
			name: "saying hello to people",
			input: input{"ZhangYundi", ""},
			want: "Hello, ZhangYundi",
		},
		{
			name: "say hello world when an empty string is supplied",
			input: input{"", "English"},
			want: "Hello, World",
		},
		{
			name: "saying hello french 1",
			input: input{"", "French"},
			want: "Bonjour, World",
		},
		{
			name: "saying hello french 2",
			input: input{"ZhangYundi", "French"},
			want: "Bonjour, ZhangYundi",
		},
		{
			name: "saying hello Spanish 1",
			input: input{"", "Spanish"},
			want: "Hola, World",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Hello(test.input.name, test.input.language)
			//assertCorrectMessage(t, got, test.want)
			assertT.Equal(test.want, got)
		})
	}
}