package whatsup

import (
	"io"
	"net/http"
	"testing"
)

func simpleHttp(url string) (resp *http.Response, err error) {
	return http.Get(url)
}

// TestSimpleHttp simply checks the response body of a url. This is a basic test to see that checking the body is a viable way of seeing whats up.
func TestIsUp(t *testing.T) {
	resp, err := simpleHttp("https://google.com")
	if err != nil {
		t.Errorf("simpleHttp failed: %v", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Failed to read response body: %v", err)
		return
	}

	if len(body) == 0 {
		t.Errorf("Response body is empty")
	}
}

// TestIsDown simply checks the response body of a url. This is a basic test to see that checking the body is a viable way of seeing whats down.
func TestIsDown(t *testing.T) {
	// sub in a url that is known to be down
	resp, err := simpleHttp("http://svc2.fourdimsdev.io/")
	if err == nil {
		t.Fatalf("testIsDown failed: %v", err)
	}

	if resp == nil {
		t.Skip("Response is nil as expected, exiting test")
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	if err == nil {
		t.Fatalf("Failed to read response body: %v", err)
		return
	}

	if len(body) != 0 {
		t.Fatalf("Response body is not empty")
	}
}
