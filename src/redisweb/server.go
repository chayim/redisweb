package redisweb

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var handler http.Handler
var srv *http.Server
var wg sync.WaitGroup = sync.WaitGroup{}

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

	wg.Add(1)
	go func() {
		fmt.Println(srv.ListenAndServe())
		wg.Done()
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

	// stop the webserver if exists, by waiting for the goroutine to exit
	if srv != nil {
		log.Printf("Stopping webserver")
		srv.Close()
		wg.Wait()
		srv = nil
	}
}

func Restart(port int) {
	Stop()
	Start(port)
}
