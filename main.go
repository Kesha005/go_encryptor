package main

import (
	"errors"
	"flag"
	"log"
	"net"
)

var addr = flag.String("listen", "172.16.11.117:8080", "listen address")

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	s := Server{Addr: *addr}
	err := s.ListenAndServe()
	if err != nil && errors.Is(err, net.ErrClosed) {
		log.Print(err)
	}
}
