package models


import (
"fmt"
"github.com/qiniu/qmgo/options"
"go.mongodb.org/mongo-driver/bson"
)

const (
	CollectionNameNft = "sync_nft"
)

type (
	NFT struct {
		Id                string  `bson:"id"`
		NTFName           string  `bson:"name"`
		DenomId           string  `bson:"denomId"`
		URI               string  `bson:"uri"`
		Data              string  `bson:"data"`
		Sender            string  `bson:"sender"`
		Recipient         string  `bson:"recipient"`
		UriHash           string  `bson:"uriHash"`
		Time     	      int64   `bson:"time"`
	}
)

func (d NFT) Name() string {
	if GetSrvConf().ChainId == "" {
		return CollectionNameNft
	}
	return fmt.Sprintf("sync_%v_nft", GetSrvConf().ChainId)
}

func (d NFT) EnsureIndexes() {
	var indexes []options.IndexModel
	indexes = append(indexes, options.IndexModel{
		Key:        []string{"-name"},
		Unique:     true,
		Background: true,
	})
	ensureIndexes(d.Name(), indexes)
}

func (d NFT) PkKvPair() map[string]interface{} {
	return bson.M{"name": d.NTFName}
}
