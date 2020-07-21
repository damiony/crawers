package db

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/crawers/go-crawers/zhenaiwang/save/model"
	"github.com/olivere/elastic/v7"
)

var (
	IndexName = "zhenaiwang"
	Type      = "user"
	es        *elastic.Client
)

type ElasticDB struct {
	Client *elastic.Client
}

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
func InsertDataToES(user *model.Profile) error {
	if es == nil {
		return errors.New("Cannot get elasticsearch client")
	}
	_, err := es.Index().
		Index(IndexName).
		Type(Type).
		Id(strconv.Itoa(user.MemberId)).
		BodyJson(user).
		Do(context.Background())
	return err
}
