package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	counter int
	port    = 9000
)

func count(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	counter++
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("{'Message': 'Hello DevFest Burgos 2017'}"))
}

func stats(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{'Visits': %d}", counter)))
}

func main() {
	//http.HandleFunc("/favicon.ico", func(_ http.ResponseWriter, _ *http.Request) {})
	http.HandleFunc("/", count)
	http.HandleFunc("/stats", stats)

	log.Println("Listening at port ", port)
	log.Panic(
		http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
