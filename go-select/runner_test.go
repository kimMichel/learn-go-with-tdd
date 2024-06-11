package goselect

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	
	slowerServer := createServerWithDelay(20 * time.Millisecond)
	fastServer := createServerWithDelay(0)

	SlowUrl := slowerServer.URL
	FastUrl := fastServer.URL

	expect := FastUrl
	result := Runner(SlowUrl, FastUrl)

	if result != expect {
		t.Errorf("expect: %s, result: %s", expect, result)
	}

	slowerServer.Close()
	fastServer.Close()
}

func createServerWithDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}