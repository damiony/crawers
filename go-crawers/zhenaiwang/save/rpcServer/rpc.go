package rpcServer

import (
	"net/rpc"

	"github.com/crawers/go-crawers/zhenaiwang/save/data"
	"github.com/crawers/go-crawers/zhenaiwang/save/model"
)

type ItemSaveService struct{}

func (itemSave *ItemSaveService) SaveItems(user *model.Profile, reply *string) error {
	err := data.SaveDataToES(user)
	if err == nil {
		*reply = "ok"
		return nil
	}
	*reply = "error"
	return err

}

func RegisterItemSaveService() error {
	return rpc.RegisterName("ItemSave", new(ItemSaveService))
}
