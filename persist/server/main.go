package main

import (
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"spider_dist/config"
	"spider_dist/persist"
	"spider_dist/rpcsupport"
)

func main() {
	log.Fatal(serveRpc(
		fmt.Sprintf(":%d", config.ItemSaverPort),
		config.ESIndex,
	))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
