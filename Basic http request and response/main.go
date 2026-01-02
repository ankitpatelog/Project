package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var startTime = time.Now()

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/status", statusHandler)

	http.ListenAndServe(":8080", nil)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format(time.RFC3339)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Time handler running  ")
	json.NewEncoder(w).Encode(map[string]string{"time": now})
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime).String()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"app":     "basic-go-server",
		"status":  "running",
		"uptime":  uptime,
		"version": "1.0.0",
	})
}
