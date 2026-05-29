package core

import (
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

var asnDB *geoip2.Reader

func InitGeoIP() {
	var err error
	asnDB, err = geoip2.Open("geo/GeoLite2-ASN.mmdb")
	if err != nil {
		log.Println("error opening geo db:", err)
	}
}

func CloseGeoIP() {
	if asnDB != nil {
		asnDB.Close()
	}
}

func LookupASN(ipStr string) (uint, string) {
	if asnDB == nil {
		return 0, "Unknown"
	}
	ip := net.ParseIP(ipStr)
	record, err := asnDB.ASN(ip)
	if err != nil {
		return 0, "Unknown"
	}
	return record.AutonomousSystemNumber, record.AutonomousSystemOrganization
}
