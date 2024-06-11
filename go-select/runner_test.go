package goselect

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)
func TestRunner(t *testing.T) {
	
	t.Run("test which url is more faster", func(t *testing.T) {
		slowerServer := createServerWithDelay(20 * time.Millisecond)
		fastServer := createServerWithDelay(0)

		SlowUrl := slowerServer.URL
		FastUrl := fastServer.URL

		expect := FastUrl
		result, _ := Runner(SlowUrl, FastUrl)

		if result != expect {
			t.Errorf("expect: %s, result: %s", expect, result)
		}

		slowerServer.Close()
		fastServer.Close()
	})

	t.Run("return an error if the server not return a response in 10s", func(t *testing.T) {
		server := createServerWithDelay(25 * time.Millisecond)

		defer server.Close()

		_, err := Configurable(server.URL, server.URL, 20 * time.Millisecond)

		if err == nil {
			t.Error("expect an error, but does not have")
		}
	})
}

func createServerWithDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}