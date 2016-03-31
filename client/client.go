package main

import (
	"flag"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"time"

	_ "github.com/gorilla/websocket"
)

var arrivalrate = flag.Int("arrivalrate", 150, "connections per second")
var maxusers = flag.Int("maxusers", 20000, "maximum number of connections to establish")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "localhost:4000", Path: "/socket/websocket"}
	log.Printf("connecting to %s -- arrivalrate = %d -- maxusers = %d\n", u.String(), *arrivalrate, *maxusers)

  count := 0
  done := make(chan struct{})
  for now := range time.Tick(time.Second / time.Duration(*arrivalrate)) {
    count += 1
    go launchClient(done, u, count, now)
    if count >= *maxusers {
    	break
    }
  }

  for count > 0 {
  	<-done
  	count--
  }
}

func launchClient(done chan struct{}, u url.URL, i int, now time.Time) {
	conn, err := net.Dial("tcp", "localhost:4000")
	log.Printf("Connect time (%d): %v\n", i, time.Since(now))
  if err != nil {
    log.Fatal(err)
  }
  defer conn.Close()

	time.Sleep(60 * time.Minute)
}