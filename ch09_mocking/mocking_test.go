package ch09_mocking

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
	"time"
)

/*小思考：mocking在这里是作为sleep 1s的模拟，通过计数的方式
这里的DI思想是剥离出sleep依赖，作为接口单独实现功能，该接口实现的方式是实现一个Sleep函数
*/

const (
	countdownStart = 3
	finalWord      = "Go!"
	write          = "write"
	sleep          = "sleep"
)

type CountdownOperationSpy struct {
	Calls []string
}

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func TestCountdown(t *testing.T) {
	assert := assert.New(t)
	subTest1 := "print 3 to Go!"
	subTest2 := "sleep before every print"
	t.Run(subTest1, func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationSpy{})
		got := buffer.String()
		want := `3
2
1
Go!`
		assert.Equal(want, got)
	})
	t.Run(subTest2, func(t *testing.T) {
		spySleepPrinter := &CountdownOperationSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		assert.Equal(want, spySleepPrinter.Calls)
	})
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()
	assert.Equal(t, sleepTime, spyTime.durationSlept)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
