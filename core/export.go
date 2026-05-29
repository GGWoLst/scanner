package core

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"sync"
)

var ExportedResults []Result
var Mu sync.Mutex

func ExportResults(results <-chan Result) {
	fCSV, _ := os.Create("results.csv")
	defer fCSV.Close()

	w := csv.NewWriter(fCSV)
	w.Write([]string{"IP", "ISP", "ASN", "PingResult", "TLSResult", "ResponseTime", "ErrorType"})

	fJSON, _ := os.Create("results.json")
	defer fJSON.Close()

	for res := range results {
		w.Write([]string{
			res.IP,
			res.ISP,
			strconv.Itoa(int(res.ASN)),
			res.PingResult,
			res.TLSResult,
			strconv.FormatInt(res.ResponseTime, 10),
			res.ErrorType,
		})
		w.Flush()

		Mu.Lock()
		ExportedResults = append(ExportedResults, res)
		data, _ := json.MarshalIndent(ExportedResults, "", "  ")
		fJSON.Seek(0, 0)
		fJSON.Truncate(0)
		fJSON.Write(data)
		Mu.Unlock()
	}
}
