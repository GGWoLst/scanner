package web

import (
	"embed"
	"encoding/json"
	"net/http"
	"scanner/core"
)

//go:embed index.html
var content embed.FS

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, _ := content.ReadFile("index.html")
		w.Write(data)
	})

	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		core.Mu.Lock()
		json.NewEncoder(w).Encode(map[string]interface{}{
			"stopped": core.StopScan,
			"count":   len(core.ExportedResults),
		})
		core.Mu.Unlock()
	})

	http.Handle("/results.csv", http.FileServer(http.Dir(".")))
	http.Handle("/results.json", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":8080", nil)
}
