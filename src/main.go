package main

import (
	"flag"
	"lgmontenegro/collantz/internal/server"
)

func main() {

	var serverPort int

	flag.IntVar(&serverPort, "p", 8080, "Specify the server port number. Default is 8080")
	
	server.Server(serverPort)	
}