package main

import (
	"fmt"
	"spider_dist/config"
	"spider_dist/rpcsupport"
	"spider_dist/worker"
	"testing"
	"time"
)

func TestSpiderService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, &worker.SpiderService{})

	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	r := worker.Request{
		URL: "http://album.zhenai.com/u/1772171421",
		Parser: &worker.SerializedParser{
			Name: config.ProfileParser,
			Args: "真爱难寻",
		},
	}

	var result worker.ParseResult
	err = client.Call(config.SpiderServiceRpc, &r, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%#v\n", result)
	}
}
