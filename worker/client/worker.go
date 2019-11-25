package client

import (
	"net/rpc"
	"spider/engine"
	"spider_dist/config"
	"spider_dist/worker"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(r *engine.Request) (*engine.ParseResult, error) {
		sr := worker.SerializeRequest(r)

		var result worker.ParseResult

		c := <-clientChan
		err := c.Call(config.SpiderServiceRpc, sr, &result)
		if err != nil {
			return &engine.ParseResult{}, err
		} else {
			return worker.DeserializeResult(&result), nil
		}
	}
}
