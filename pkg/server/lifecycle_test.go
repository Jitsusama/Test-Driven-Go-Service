package server_test

import (
	"fmt"
	"net"
	"playing-around/pkg/server"
	"testing"
	"time"
)

func TestListens(t *testing.T) {
	for _, port := range []int{freeTcpPort(), freeTcpPort()} {
		t.Run(fmt.Sprintf("on port %d", port), func(t *testing.T) {
			// Arrange
			svr := server.Create(port, nil)
			go func() { _ = svr.Start() }()
			defer func() { _ = svr.Stop() }()
			time.Sleep(time.Millisecond)

			// Act
			conn, err := net.Dial("tcp4", fmt.Sprintf(":%d", port))
			if err == nil {
				defer func() { _ = conn.Close() }()
			}

			// Assert
			if err != nil {
				t.Fatalf("unable to connect to port %d: %v", port, err)
			}
		})
	}
}

func TestStopsListening(t *testing.T) {
	for _, port := range []int{freeTcpPort(), freeTcpPort()} {
		t.Run(fmt.Sprintf("on port %d", port), func(t *testing.T) {
			// Arrange
			svr := server.Create(port, nil)
			go func() { _ = svr.Start() }()
			time.Sleep(time.Millisecond)
			_ = svr.Stop()

			// Act
			conn, err := net.DialTimeout("tcp4", fmt.Sprintf(":%d", port), time.Millisecond*10)
			if err == nil {
				defer func() { _ = conn.Close() }()
			}

			// Assert
			if err == nil {
				t.Fatalf("server is still listening on port %d", port)
			}
		})
	}
}

func freeTcpPort() int {
	l, _ := net.Listen("tcp4", ":0")
	defer func(l net.Listener) { _ = l.Close() }(l)
	return l.Addr().(*net.TCPAddr).Port
}
