package whatsup

import (
	"net/http"
	"testing"
)

// TestSimpleHttp simply checks the response body of a url. This is a basic test to see that checking the body is a viable way of seeing whats up.
func TestIsUp(t *testing.T) {
	up := IsResponseLive(http.Get("https://google.com"))
	if !up {
		t.Errorf("Expected response to be up")
	}
}

// TestIsDown simply checks the response body of a url. This is a basic test to see that checking the body is a viable way of seeing whats down.
func TestIsDown(t *testing.T) {
	// sub in a url that is known to be down
	up := IsResponseLive(http.Get("http://svc2.fourdimsdev.io/"))
	if up {
		t.Errorf("Expected response to be down")
	}
}
