package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	// Read arg from command line for the port.
	portPtr := flag.String("port", "9000", "a string")
	flag.Parse()
	log.Printf("API Demo is running in port: %q", *portPtr)
	apiPort := string(*portPtr)
	addr := fmt.Sprintf(":%v", apiPort)

	api := App{}
	api.Initialize("", "", "")
	api.Run(addr)
}
