package handlers

import (
	"github.com/FiiLabs/block_explorer/config"
	"github.com/FiiLabs/block_explorer/libs/pool"
	"github.com/FiiLabs/block_explorer/models"
	"github.com/FiiLabs/block_explorer/utils"
	. "github.com/kaifei-bianjie/tibc-mod-parser/modules"
	"testing"
)

func TestParseTxs(t *testing.T) {
	block := int64(3744282)
	conf, err := config.ReadConfig()
	if err != nil {
		t.Fatal(err.Error())
	}
	models.Init(conf)
	InitRouter(conf)
	pool.Init(conf)
	c := pool.GetClient()
	defer func() {
		c.Release()
	}()

	if blockDoc, txDocs, NFTClses, NFTs, err := ParseBlockAndTxs(block, c); err != nil {
		t.Fatal(err)
	} else {
		t.Log(utils.MarshalJsonIgnoreErr(blockDoc))
		t.Log(utils.MarshalJsonIgnoreErr(txDocs))
		t.Log(utils.MarshalJsonIgnoreErr(NFTClses))
		t.Log(utils.MarshalJsonIgnoreErr(NFTs))
	}
}

func TestUnmarshalTibcAckInEvents(t *testing.T) {
	bytesdata := []byte("\ufffd\u0001\u0001\u0001")
	var result TIBCAcknowledgement
	err := result.Unmarshal(bytesdata)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(result.String())
}
