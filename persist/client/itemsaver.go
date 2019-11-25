package client

import (
	"log"
	"spider/engine"
	"spider_dist/config"
	"spider_dist/rpcsupport"
)

func ItemSaver(host string) (chan *engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	ch := make(chan *engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-ch
			log.Printf("Got item#%d: %+v", itemCount, item)
			itemCount++

			// Call RPC to save item
			var result string
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item Saver: error saving item: %v: %v", item, err)
			}
		}
	}()

	return ch, nil
}
