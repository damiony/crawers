package persist

import (
	"crawer/engine"
	"fmt"
	"net/rpc"
	"time"
)

func ItemSave(client *rpc.Client) chan engine.Item {
	out := make(chan engine.Item)
	go Save(out, client)

	return out
}

func Save(out chan engine.Item, client *rpc.Client) {
	var reply string
	for {
		time.Sleep(50 * time.Millisecond)
		item := <-out
		go func() {
			err := client.Call("ItemSaveService.SaveItems", item, &reply)
			if err != nil {
				fmt.Printf("Save items error: %v\n", err)
			}
		}()
	}

}
