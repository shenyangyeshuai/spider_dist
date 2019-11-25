package main

import (
	"flag"
	"fmt"
	"log"
	"spider_dist/rpcsupport"
	"spider_dist/worker"
)

var (
	port = flag.Int("port", 0, "the port for me to listen on")
)

func main() {
	flag.Parse()

	if *port == 0 {
		fmt.Println("必须指定 port")
		return
	}

	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		&worker.SpiderService{},
	))
}
