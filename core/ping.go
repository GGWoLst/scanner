package core

import (
	"time"

	"github.com/go-ping/ping"
)

func PingNode(ip string) (string, int64) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		return "Error", 0
	}
	pinger.Count = 1
	pinger.Timeout = 2 * time.Second
	pinger.SetPrivileged(true)

	err = pinger.Run()
	if err != nil {
		return "Error", 0
	}

	stats := pinger.Statistics()
	if stats.PacketsRecv > 0 {
		return "Success", stats.AvgRtt.Milliseconds()
	}
	return "Timeout", 0
}
