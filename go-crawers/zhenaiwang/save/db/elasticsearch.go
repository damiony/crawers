package db

import (
	"context"
	"errors"
	"log"
	"save/engine"
	"strconv"

	"github.com/olivere/elastic/v7"
)

var (
	IndexName = "zhenaiwang"
	Type      = "user"
	es        *elastic.Client
)

// initialize
func init() {
	if es != nil {
		return
	}
	options := []elastic.ClientOptionFunc{
		elastic.SetSniff(false),
	}
	client, err := elastic.NewClient(options...)
	if err != nil {
		log.Fatal(err)
	}
	es = client
	return
}

// insert data in elasticsearch
func InsertDataToES(item *engine.Item) error {
	if es == nil {
		return errors.New("Cannot get elasticsearch client")
	}
	if item.Type == "" || item.Id == 0 {
		return errors.New("Invalid type or id")
	}
	_, err := es.Index().
		Index(IndexName).
		Type(item.Type).
		Id(strconv.Itoa(item.Id)).
		BodyJson(item).
		Do(context.Background())
	return err
}
