package main

import (
	"os"
	"github.com/zz9629/cloudgo/service"
	flag "github.com/spf13/pflag"
)

// Default cloudgo network port 8080
const (
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	// use pflag to get the parameter -p
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	server := service.NewServer()
	server.Run(":" + port)
}