package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"spider/engine"
	"spider/scheduler"
	"spider/zhenai/parser"
	"spider_dist/config"
	is "spider_dist/persist/client"
	"spider_dist/rpcsupport"
	wr "spider_dist/worker/client"
	"strings"
)

var (
	workerHosts = flag.String(
		"worker_hosts",
		"",
		"worker hosts (comma separated)",
	)
)

func main() {
	flag.Parse()

	// Concurrent Engine
	itemChan, err := is.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	hs := strings.Split(*workerHosts, ",")
	pool := createClientPool(hs)
	processor := wr.CreateProcessor(pool)

	concurrentEngine := &engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	concurrentEngine.Run(&engine.Request{
		URL: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList,
		),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	c := make(chan *rpc.Client)

	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("connected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}
	}

	go func() {
		for {
			for _, client := range clients {
				c <- client
			}
		}
	}()

	return c
}
