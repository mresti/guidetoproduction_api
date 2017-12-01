package main

import (
	"flag"
	"fmt"
	"log"
	"api/app"
)

func main() {
	// Read arg from command line for the port.
	portPtr := flag.String("port", "9000", "a string")
	flag.Parse()
	log.Printf("API Demo is running in port: %q", *portPtr)
	apiPort := string(*portPtr)
	addr := fmt.Sprintf(":%v", apiPort)

	api := app.App{}
	api.Initialize("", "", "")
	api.Run(addr)
}
