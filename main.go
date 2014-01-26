package main

import (
	"flag"
	"log"
	"net"
	"net/http/fcgi"
)

func main() {
	m := NewMartiniServer()

	var use_fcgi bool
	flag.BoolVar(&use_fcgi, "fcgi", false, "Hosts with FCGI on port 9000")
	flag.Parse()

	if use_fcgi {
		log.Print("Serving via FCGI...")
		listener, err := net.Listen("tcp", "127.0.0.1:9000")
		if err != nil {
			log.Fatal(err)
		}
		fcgi.Serve(listener, m)
	} else {
		log.Print("Serving via straight Martini...")
		m.Run()
	}
}
