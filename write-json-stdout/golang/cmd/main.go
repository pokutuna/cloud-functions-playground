package main

import (
	"net/http"

	function "github.com/pokutuna/cloud-functions-playground/write-json-stdout/golang"
)

func main() {
	http.HandleFunc("/", function.App)
	http.ListenAndServe(":3000", nil)
}
