package rpcClient

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"

	"github.com/go-crawler/zhenaiwang/model"
)

func ItemSave(itemChan chan model.Profile, host string) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}

	client := jsonrpc.NewClient(conn)

	var reply string
	for {
		item := <-itemChan
		go func() {
			err := client.Call("ItemSaveService.SaveItems", item, &reply)
			if err != nil {
				fmt.Printf("Save items error: %v", err)
			}
		}()
	}
}
