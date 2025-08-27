package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

type JSONResponse struct {
	Message   string     `json:"message"`
	Timestamp int64      `json:"timestamp"`
	Data      []DataItem `json:"data"`
}

type DataItem struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type CPUResponse struct {
	Input  int     `json:"input"`
	Result int64   `json:"result"`
	TimeMs float64 `json:"time_ms"`
}

type PostDataResponse struct {
	ReceivedCount int   `json:"received_count"`
	ProcessedAt   int64 `json:"processed_at"`
	Summary       int   `json:"summary"`
}

func fibonacci(n int) int64 {
	if n <= 1 {
		return int64(n)
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().UnixMilli(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Hello World!")
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	data := make([]DataItem, 100)
	for i := 0; i < 100; i++ {
		data[i] = DataItem{
			ID:    i,
			Value: fmt.Sprintf("item-%d", i),
		}
	}

	response := JSONResponse{
		Message:   "Hello World!",
		Timestamp: time.Now().UnixMilli(),
		Data:      data,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func cpuHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nStr, exists := vars["n"]
	if !exists {
		nStr = "40"
	}

	n, err := strconv.Atoi(nStr)
	if err != nil {
		n = 40
	}

	start := time.Now()
	result := fibonacci(n)
	elapsed := time.Since(start)

	response := CPUResponse{
		Input:  n,
		Result: result,
		TimeMs: float64(elapsed.Nanoseconds()) / 1000000.0,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func postDataHandler(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	jsonBytes, _ := json.Marshal(data)

	var receivedCount int
	switch v := data.(type) {
	case []interface{}:
		receivedCount = len(v)
	case map[string]interface{}:
		receivedCount = len(v)
	default:
		receivedCount = 1
	}

	response := PostDataResponse{
		ReceivedCount: receivedCount,
		ProcessedAt:   time.Now().UnixMilli(),
		Summary:       len(jsonBytes),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler).Methods("GET")
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/json", jsonHandler).Methods("GET")
	r.HandleFunc("/cpu/{n:[0-9]+}", cpuHandler).Methods("GET")
	r.HandleFunc("/data", postDataHandler).Methods("POST")

	fmt.Println("Go HTTP server running on port 3001")
	log.Fatal(http.ListenAndServe(":3001", r))
}
