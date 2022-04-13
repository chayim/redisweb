package main

import "C"

import (
	"github.com/chayim/redisweb/src/redisweb"
)

//export HTTPStart
func HTTPStart(p *C.int) {
	port := int(*p)
	redisweb.Start(port)
}

//export HTTPStop
func HTTPStop() {
	redisweb.Stop()
}

//export HTTPRestart
func HTTPRestart(p *C.int) {
	port := int(*p)
	redisweb.Restart(port)
}

func main() {}
