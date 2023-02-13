package main

import (
	"fmt"
	"net/http"
	"os"

	"example.com/function"
)

func main() {
	port := "3000"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	http.HandleFunc("/", function.Function)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
