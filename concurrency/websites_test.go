package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
        "http://blog.gypsydave5.com",
        "waat://furhurterwe.geds",
	}

	expect := map[string]bool {
		"http://google.com":          true,
        "http://blog.gypsydave5.com": true,
        "waat://furhurterwe.geds":    false,
	}

	result := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("expect: %v, result: %v", expect, result)
	}
}