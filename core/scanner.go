package core

import (
	"sync"
)

var StopScan bool

func StartScan(ips []string, results chan<- Result) {
	sem := make(chan struct{}, 1000)
	var wg sync.WaitGroup

	for _, ip := range ips {
		if StopScan {
			break
		}
		wg.Add(1)
		sem <- struct{}{}
		
		go func(target string) {
			defer wg.Done()
			defer func() { <-sem }()
			
			if StopScan {
				return
			}

			pingRes, rtt := PingNode(target)
			tlsRes, errType := CheckTLS(target)
			asn, isp := LookupASN(target)

			results <- Result{
				IP:           target,
				ISP:          isp,
				ASN:          asn,
				PingResult:   pingRes,
				TLSResult:    tlsRes,
				ResponseTime: rtt,
				ErrorType:    errType,
			}
		}(ip)
	}

	wg.Wait()
	close(results)
}
