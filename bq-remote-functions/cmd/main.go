package main

import (
	"fmt"
	"net/http"
	"os"

	function "github.com/pokutuna/cloud-functions-playground/bq-remote-functions"
)

func main() {
	port := "3000"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	http.HandleFunc("/", function.RemoteFunction)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
