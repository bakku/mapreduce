package main

import (
	"flag"
	"fmt"
	"log"
	"mapreduce"
	"net"
	"net/http"
	"net/rpc"
)

var port = flag.Int("port", 1234, "the port of the RPC server")

func main() {
	flag.Parse()

	m := &mapreduce.JobScheduler{Workers: []mapreduce.Worker{}}

	err := rpc.Register(m)
	if err != nil {
		log.Fatalln("error registering rpc job scheduler:", err)
	}

	rpc.HandleHTTP()

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalln("error listening on port:", err)
	}

	if err := http.Serve(l, nil); err != nil {
		log.Fatal(err)
	}
}
