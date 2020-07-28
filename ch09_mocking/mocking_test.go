package ch09_mocking

//小思考：mocking在这里是作为sleep 1s的模拟，通过计数的方式
//这里的DI思想是剥离出sleep依赖，作为接口单独实现功能，该接口实现的方式是实现一个Sleep函数

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

const (
	finalWorld     = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

// mock: 可以记录依赖关系是怎样被使用的，可以用来记录传入进来的参数 多少次。。。在这里记录了Sleep()被调用多少次
func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWorld)
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	Countdown(buffer, spySleeper)
	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}
