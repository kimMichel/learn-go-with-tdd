package goselect

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	
	slowerServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

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