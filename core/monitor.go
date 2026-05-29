package core

import (
	"log"
	"time"
)

func MonitorConnectivity() {
	for {
		res, _ := PingNode("1.1.1.1")
		if res == "Success" {
			log.Println("Whitelist is disabled (1.1.1.1 is accessible). Stopping scan.")
			StopScan = true
			return
		}
		time.Sleep(10 * time.Second)
	}
}
