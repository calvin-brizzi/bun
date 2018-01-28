// Package bun creates registers a "/search" endpoint that acts as a smart redirector
// "/" gives a list of current commands
package bun

import (
	"flag"
	"net/http"
)

var serverAddr = flag.String("addr", ":8080", "Address of the server")

func init() {
	flag.Parse()

	initCustom()
	http.HandleFunc("/", helpHandler)
	http.HandleFunc("/search", searchHandler)
}
