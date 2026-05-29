package core

import (
	"crypto/tls"
	"net"
	"strings"
	"syscall"
	"time"
)

func CheckTLS(ip string) (string, string) {
	dialer := &net.Dialer{Timeout: 3 * time.Second}
	conf := &tls.Config{InsecureSkipVerify: true}

	conn, err := tls.DialWithDialer(dialer, "tcp", net.JoinHostPort(ip, "443"), conf)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			return "Timeout", "Drop"
		}
		errStr := err.Error()
		if strings.Contains(errStr, "connection reset by peer") || strings.Contains(errStr, syscall.ECONNRESET.Error()) {
			return "Error", "TCP RST (DPI)"
		}
		if strings.Contains(errStr, "connection refused") {
			return "Error", "Connection Refused"
		}
		return "Error", "Handshake Failed"
	}
	conn.Close()
	return "Success", ""
}
