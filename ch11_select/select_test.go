package ch11_select

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
		}))
	fastServer := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)
	assert.Equal(t, want, got)
	slowServer.Close()
	fastServer.Close()
}

func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)
	if aDuration < bDuration {
		return a
	}
	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
