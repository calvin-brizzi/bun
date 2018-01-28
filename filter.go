package bun

import (
	"flag"
	"log"
	"net/http"
)

var useIPWhitelisting = flag.Bool("whitelist_ips", false, "Whether or not to use IP Whitelisting")

// The ipWhitelist for custom functions you don't want to
// expose to the larger world
var ipWhitelist = map[string]bool{
	"127.0.0.1":      true,
}

func shouldFilter(r *http.Request) bool {
	if !*useIPWhitelisting {
		return false
	}
	key := r.Header.Get("X-Real-IP")
	log.Printf(key)
	return !(ipWhitelist[key] || key == "")
}

func sendToDefault(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, Commands[defaultCommand].redirectFunc(""), 301)
}
