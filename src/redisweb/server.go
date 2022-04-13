package redisweb

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var handler http.Handler
var srv *http.Server

func Start(port int) bool {

	if srv != nil {
		return false
	}

	// default 8080
	if port < 1 || port > 65536 {
		port = 8080
	}
	handler = http.HandlerFunc(handleRequest)
	log.Println("Starting web server")
	srv = &http.Server{Addr: fmt.Sprintf(":%d", port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  12 * time.Second,
		Handler:      handler,
	}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	return true
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	u, p, ok := r.BasicAuth()
	if !ok {
		fmt.Println("Error parsing basic auth")
		w.WriteHeader(401)
		return
	}

	// TODO auth here
	fmt.Println(u)
	fmt.Println(p)

	w.WriteHeader(200)
	return
}

func Stop() {
	if srv != nil {
		log.Printf("Stopping webserver")
		srv.Close()
	}
}

func Restart(port int) {
	Stop()
	Start(port)
}
