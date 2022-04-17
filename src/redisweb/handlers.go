package redisweb

import (
	"fmt"
	"net/http"
)

func authenticate(w http.ResponseWriter, r *http.Request) bool {
	return basicAuth(w, r)
}

func basicAuth(w http.ResponseWriter, r *http.Request) bool {
	u, p, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(401)
		w.Write([]byte("Authentication failed."))
		return false
	}

	// TODO handle auth here
	fmt.Println(u)
	fmt.Println(p)

	return true
}

func pipelineHandler(w http.ResponseWriter, r *http.Request) {
	if !authenticate(w, r) {
		return
	}
}

func syncHandler(w http.ResponseWriter, r *http.Request) {
	if !authenticate(w, r) {
		return
	}

	parts, err := despeckle(r.URL.RawQuery)
	if err != nil {
		fmt.Println("Failed to parse query string")
		w.WriteHeader(401)
		return
	}
	Execute(parts)
}
