package redisweb

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

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

	log.Println(fmt.Sprintf("Starting web server on port %d", port))
	srv = &http.Server{Addr: fmt.Sprintf(":%d", port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  12 * time.Second,
	}

	// default is sync
	http.HandleFunc("/", syncHandler)
	http.HandleFunc("/pipeline", pipelineHandler)
	http.HandleFunc("/sync", syncHandler)

	wg.Add(1)
	go func() {
		fmt.Println(srv.ListenAndServe())
		wg.Done()
	}()
	return true
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
