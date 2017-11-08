package main

import (
	"fmt"
	"log"
	"net/http"
	"flag"
)

var (
	counter int
)

func count(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	counter++
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("{'Message': 'Hello DevFest Granada 2017'}"))
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
	// Read arg from command line for the port.
	portPtr := flag.String("port", "9000", "a string")
	flag.Parse()
	log.Printf("API Demo is running in port: %q", *portPtr)

	//http.HandleFunc("/favicon.ico", func(_ http.ResponseWriter, _ *http.Request) {})
	http.HandleFunc("/", count)
	http.HandleFunc("/stats", stats)

	apiPort := string(*portPtr)

	addr := fmt.Sprintf(":%v", apiPort)
	log.Fatal(http.ListenAndServe(addr, nil))
}
