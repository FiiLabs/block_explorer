package models


import (
"fmt"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionNameNft = "sync_nft"
)
type ActionNFT int
const (
	TxActionNFTMint ActionNFT = iota
	TxActionNFTBurn
	TxActionNFTEdit
	TxActionNFTTransfer
)
type (

	NFT struct {
		iD             primitive.ObjectID `bson:"_id"`
		Id                string  `bson:"id"`
		NTFName           string  `bson:"name"`
		DenomId           string  `bson:"denomId"`
		URI               string  `bson:"uri"`
		Data              string  `bson:"data"`
		Sender            string  `bson:"sender"`
		Recipient         string  `bson:"recipient"`
		UriHash           string  `bson:"uriHash"`
		Time     	      int64   `bson:"time"`
		action            ActionNFT
	}
	NFTWrap struct {
		ID             primitive.ObjectID `bson:"_id"`
		Id                string  `bson:"id"`
		NTFName           string  `bson:"name"`
		DenomId           string  `bson:"denomId"`
		URI               string  `bson:"uri"`
		Data              string  `bson:"data"`
		Sender            string  `bson:"sender"`
		Recipient         string  `bson:"recipient"`
		UriHash           string  `bson:"uriHash"`
		Time     	      int64   `bson:"time"`
		action            ActionNFT
	}
)
func (d *NFT) SetAction(action ActionNFT) {
	d.action = action
}
func (d *NFT) GetAction() ActionNFT{
	return d.action
}
func (d NFT) Name() string {
	if GetSrvConf().ChainId == "" {
		return CollectionNameNft
	}
	return fmt.Sprintf("sync_%v_nft", GetSrvConf().ChainId)
}

func (d NFT) EnsureIndexes() {
	var indexes []options.IndexModel
	indexes = append(indexes, options.IndexModel{
		Key:        []string{"-id"},
		Unique:     true,
		Background: true,
	})
	ensureIndexes(d.Name(), indexes)
}

func (d NFT) PkKvPair() map[string]interface{} {
	return bson.M{"id": d.Id}
}

func (d NFT) GetNFTById(id string) (primitive.ObjectID, error) {
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

func (d *NFT) SetID(id primitive.ObjectID) {
	//d.iD = id
	copy(d.iD[:], id[:])
}

func (d *NFT) GetID( ) primitive.ObjectID{
	return d.iD
}