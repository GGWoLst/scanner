package core

type Result struct {
	IP           string `json:"ip"`
	ISP          string `json:"isp"`
	ASN          uint   `json:"asn"`
	PingResult   string `json:"ping_result"`
	TLSResult    string `json:"tls_result"`
	ResponseTime int64  `json:"response_time"`
	ErrorType    string `json:"error_type"`
}
