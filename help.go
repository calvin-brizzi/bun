package bun

import (
	"html/template"
	"net/http"
)

// helpHandler prints the help page
func helpHandler(w http.ResponseWriter, r *http.Request) {
	if shouldFilter(r) {
		sendToDefault(w, r)
		return
	}
	t, err := template.New("help").Parse(helpTemplate)
	if err != nil {
		panic(err)
	}
	t.Execute(w, Commands)
}
