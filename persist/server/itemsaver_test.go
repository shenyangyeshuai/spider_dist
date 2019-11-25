package main

import (
	"spider/engine"
	"spider/model"
	"spider_dist/config"
	"spider_dist/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	host := ":12345"
	// start ItemSaverServer
	go serveRpc(host, "test1")

	time.Sleep(time.Second)

	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call save
	item := engine.Item{
		URL:  "https://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: model.Profile{
			Name:       "安静的雪",
			Gender:     "女",
			Age:        32,
			Height:     160,
			Income:     "3001-5000元",
			Marriage:   "已婚",
			Occupation: "人事/行政",
			Xingzuo:    "牧羊座",
		},
	}

	var result string
	err = client.Call(config.ItemSaverRpc, &item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
