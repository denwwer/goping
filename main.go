package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var errorResp = []byte(`{"status": "failed"}`)
var okResp = `{"status": "ok", "uptime": "%s"}`

func main() {
	port := flag.String("p", "45331", "Set port for /ping endpoint")
	flag.Parse()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		uptimeBytes, err := os.ReadFile("/proc/uptime")
		if err != nil {
			log.Println("[ERROR] goping uptime error:", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorResp)
			return
		}

		uptimeStr := strings.Split(string(uptimeBytes), " ")[0]
		uptimeSec, err := strconv.ParseFloat(uptimeStr, 64)
		if err != nil {
			log.Println("[ERROR] goping parsing error:", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorResp)
			return
		}

		// Format
		uptimeDur := time.Duration(uptimeSec) * time.Second
		days := int(uptimeDur.Hours() / 24)
		hours := int(uptimeDur.Hours()) % 24
		minutes := int(uptimeDur.Minutes()) % 60

		var uptime string
		if days > 0 {
			uptime = fmt.Sprintf("%d days %d:%02d", days, hours, minutes)
		} else {
			uptime = fmt.Sprintf("%d:%02d", hours, minutes)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, okResp, uptime)
	})

	log.Printf("[INFO] goping started on port: %s", *port)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal("[ERROR] goping failed to start:", err)
	}
}
