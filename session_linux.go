//go:build linux
package neugo

import (
    "net"
	"net/http"
	"net/http/cookiejar"
	"syscall"
	"time"
)

// NewSession returns a *http.Client with an empty cookie jar ,timeout of 6s and transport with specific fwmark.
func NewFwmarkSession(mark uint32) *http.Client {
	dialer := &net.Dialer{
		Control: func(network, address string, c syscall.RawConn) error {
			return c.Control(func(fd uintptr) {
				err := syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_MARK, int(mark))
				if err != nil {
					return
				}
			})
		},
	}

	tr := &http.Transport{
		DialContext: dialer.DialContext,
	}

	jar, _ := cookiejar.New(nil)
	n := &http.Client{
		Timeout:   6 * time.Second,
		Jar:       jar,
		Transport: tr,
	}
	return n
}
