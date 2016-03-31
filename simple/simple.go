package main

import (
	"flag"
	"log"
	"net"
	"time"
)

var host = flag.String("server", "localhost:4000", "server's hostname")
var arrivalrate = flag.Int("arrivalrate", 150, "connections per second")
var maxusers = flag.Int("maxusers", 20000, "maximum number of connections to establish")

func main() {
	flag.Parse()
	log.SetFlags(0)

	// Hold onto conn's so they don't get GC'ed
	sc, cc := make([]net.Conn, *maxusers), make([]net.Conn, *maxusers)

	// Start server
	log.Printf("Starting server...\n")
	ln, err := net.Listen("tcp", *host)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	// Spawn an Accept loop
	go func() {
		counter := 0
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Fatal(err)
			}
			sc[counter] = conn
			counter += 1
		}
	}()

	// Spin up connections
	log.Printf("connecting to %s -- arrivalrate = %d -- maxusers = %d\n", host, *arrivalrate, *maxusers)
	count := 0
	for now := range time.Tick(time.Second / time.Duration(*arrivalrate)) {
		cc[count] = launchClient(count, now)
		count += 1
		if count >= *maxusers {
			break
		}
	}
}

func launchClient(i int, now time.Time) net.Conn {
	conn, err := net.Dial("tcp", *host)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connect time (%d): %v\n", i, time.Since(now))
	//time.Sleep(60 * time.Minute)
	return conn
}
