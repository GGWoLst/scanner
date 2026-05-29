package main

import (
	"flag"
	"fmt"
	"log"
	"scanner/core"
	"scanner/web"
	"time"
)

func main() {
	webMode := flag.Bool("web", false, "Start with Web UI on port 8080")
	flag.Parse()

	core.InitGeoIP()
	defer core.CloseGeoIP()

	go core.MonitorConnectivity()

	ips := core.LoadIPs()
	results := make(chan core.Result, len(ips))

	go core.StartScan(ips, results)
	go core.ExportResults(results)

	if *webMode {
		log.Println("Starting in Web Mode on :8080...")
		web.StartServer()
	} else {
		log.Println("Starting in CLI Mode...")
		total := len(ips)
		for {
			core.Mu.Lock()
			scanned := len(core.ExportedResults)
			stopped := core.StopScan
			core.Mu.Unlock()

			fmt.Printf("\rScanned: %d / %d", scanned, total)

			if stopped || scanned >= total {
				fmt.Println("\nScan completed or stopped.")
				break
			}
			time.Sleep(1 * time.Second)
		}
	}
}
