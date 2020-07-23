package data

import (
	"save/db"
	"save/engine"
)

func SaveDataToES(item *engine.Item) error {
	return db.InsertDataToES(item)
}
