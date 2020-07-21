package data

import (
	"testing"

	"github.com/crawers/go-crawers/zhenaiwang/save/model"
)

func TestSaveDataToES(t *testing.T) {
	user := model.Profile{
		MemberId:  111,
		NickName:  "hahaha",
		Education: "本科",
		Age:       18,
		Marriage:  "已婚",
		Height:    175,
		Sex:       0,
	}
	err := SaveDataToES(&user)
	if err != nil {
		t.Errorf("Save user fail: %+v\n", user)
	}
}
