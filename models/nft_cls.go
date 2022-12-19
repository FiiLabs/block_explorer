package models

import (
	"fmt"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionNameNftCls = "sync_nft_cls"
)
type ActionNFTCls int
const (
	TxActionNFTClsMint ActionNFTCls = iota
	TxActionNFTClsTransfer
)
type (
	NFTCls struct {
		iD             primitive.ObjectID `bson:"_id"`
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
		action            ActionNFTCls
	}

	NFTClsWrap struct {
		ID             primitive.ObjectID `bson:"_id"`
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
		action            ActionNFTCls
	}
)
func (d *NFTCls) SetAction(action ActionNFTCls) {
	d.action = action
}
func (d *NFTCls) GetAction() ActionNFTCls{
	return d.action
}
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


func (d NFTCls) GetNFTClsById(id string) (primitive.ObjectID, error) {
	var nft NFTWrap

	fn := func(c *qmgo.Collection) error {
		return c.Find(_ctx, bson.M{
			"id": id,
		}).One(&nft)
	}

	err := ExecCollection(d.Name(), fn)
	if err != nil {
		return nft.ID, err
	}
	return nft.ID, nil
}

func (d *NFTCls) SetID(id primitive.ObjectID) {
	//d.iD = id
	copy(d.iD[:], id[:])
}

func (d *NFTCls) GetID( ) primitive.ObjectID{
	return d.iD
}