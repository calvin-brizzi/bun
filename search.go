package bun

import (
	"net/http"
	"strings"
)

type redirector func(string) string

// Command defines a command Bun accepts and describes its usage
// Name, Key and Help are strings used to render the help page
// redirectFunc is a function that maps an incoming string (the query)
// to a url (as a string) to redirect to
// The private flag determines whether the query should just be redirected to
// google if the requesting IP is not in the whitelist
type Command struct {
	Name         string
	Key          string
	redirectFunc redirector
	Help         string
	private      bool
}

// Handler is the main entrypoint to the program
func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()["q"][0]
	selector := query
	i := strings.Index(query, " ")
	if i > 0 {
		selector = query[:i]
	}

	command, ok := Commands[selector]
	if !ok || (command.private && shouldFilter(r)) {
		command = Commands[defaultCommand]
		i = -1
	}

	redirectURL := command.redirectFunc(query[i+1:])

	http.Redirect(w, r, redirectURL, 301)
}
