package persist

import (
	"fmt"
	"net/rpc"
	"time"

	"github.com/go-crawler/zhenaiwang/engine"
)

func ItemSave(client *rpc.Client) chan engine.Item {
	out := make(chan engine.Item)
	go Save(out, client)

	return out
}

func Save(out chan engine.Item, client *rpc.Client) {
	var reply string
	for {
		item := <-out
		go func() {
			err := client.Call("ItemSaveService.SaveItems", item, &reply)
			if err != nil {
				fmt.Printf("Save items error: %v\n", err)
			}
		}()
		time.Sleep(time.Second)
	}

}
