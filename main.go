package main

import (
	"fmt"
	"github.com/Lorenc326/vitl-test/api"
)

func main() {
	server := api.BuildServer()
	port := "8080"
	address := fmt.Sprintf(":%s", port)
	server.Logger.Fatal(server.Start(address))
}