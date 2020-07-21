package data

import (
	"testing"

	"github.com/crawers/go-crawers/zhenaiwang/save/model"
)

func TestSaveData(t *testing.T) {
	var user = model.Profile{
		NickName:      "test",
		Constellation: "天蝎",
		Education:     "大学",
		Age:           20,     // 年龄
		Marriage:      "未婚",   // 婚姻
		Height:        165,    // 身高
		MemberId:      123456, // ID
		Sex:           0,      // 性别 0男 1女
	}
	err := SaveDataToMysql(&user)
	if err != nil {
		t.Errorf("save error: %v", err)
	} else {
		t.Log("save suc")
	}
}
