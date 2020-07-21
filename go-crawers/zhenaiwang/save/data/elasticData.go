package data

import (
	"github.com/crawers/go-crawers/zhenaiwang/save/db"
)

func SaveDataToES(user interface{}) error {
	return db.InsertDataToES(user)
}
