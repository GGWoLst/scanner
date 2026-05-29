package core

import (
	"bufio"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func LoadIPs() []string {
	var ips []string
	
	files, err := os.ReadDir("cidr")
	if err != nil {
		return ips
	}

	for _, entry := range files {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".txt" {
			continue
		}

		file, err := os.Open(filepath.Join("cidr", entry.Name()))
		if err != nil {
			continue
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" {
				continue
			}
			
			// Проверяем, это CIDR или просто IP
			if strings.Contains(line, "/") {
				parsedIPs := expandCIDR(line)
				ips = append(ips, parsedIPs...)
			} else {
				ips = append(ips, line)
			}
		}
		file.Close()
	}

	return ips
}

func expandCIDR(cidr string) []string {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil
	}
	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	if len(ips) > 2 {
		return ips[1 : len(ips)-1]
	}
	return ips
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
