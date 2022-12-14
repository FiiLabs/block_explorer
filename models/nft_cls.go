package models

import (
	"fmt"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	CollectionNameNftCls = "sync_nft_cls"
)

type (
	NFTCls struct {
		Id                string  `bson:"id"`
		ClsName           string  `bson:"name"`
		Schema            string  `bson:"schema"`
		Sender            string  `bson:"sender"`
		Symbol            string  `bson:"symbol"`
		MintRestricted    bool    `bson:"mintRestricted"`
        UpdateRestricted  bool    `bson:"updateRestricted"`
		Description       string  `bson:"description"`
		Uri               string  `bson:"uri"`
		UriHash           string  `bson:"uriHash"`
		Data              string  `bson:"data"`
		Time     	      int64   `bson:"time"`
	}
)

func (d NFTCls) Name() string {
	if GetSrvConf().ChainId == "" {
		return CollectionNameNftCls
	}
	return fmt.Sprintf("sync_%v_nftcls", GetSrvConf().ChainId)
}

func (d NFTCls) EnsureIndexes() {
	var indexes []options.IndexModel
	indexes = append(indexes, options.IndexModel{
		Key:        []string{"-name"},
		Unique:     true,
		Background: true,
	})
	ensureIndexes(d.Name(), indexes)
}

func (d NFTCls) PkKvPair() map[string]interface{} {
	return bson.M{"name": d.ClsName}
}