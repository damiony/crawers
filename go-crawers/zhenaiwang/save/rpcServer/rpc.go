package rpcServer

import (
	"fmt"
	"net/rpc"
	"save/data"
	"save/engine"
)

type ItemSaveService struct{}

func (itemSave *ItemSaveService) SaveItems(item *engine.Item, reply *string) error {
	fmt.Printf("Save item: %+v\n", item)
	err := data.SaveDataToES(item)
	if err == nil {
		*reply = "ok"
		return nil
	}
	*reply = "error"
	return err

}

func RegisterItemSaveService() error {
	return rpc.RegisterName("ItemSaveService", new(ItemSaveService))
}
