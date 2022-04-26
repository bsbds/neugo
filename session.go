package neugo

import (
	"net/http"
	"net/http/cookiejar"
	"time"
)

// NewSession returns a *http.Client with an empty cookie jar and timeout of 6s.
func NewSession() *http.Client {
	jar, _ := cookiejar.New(nil)
	n := &http.Client{
		Timeout: 6 * time.Second,
		Jar:     jar,
	}
	return n
}
