package rest

import (
	"github.com/FiiLabs/block_explorer/service"
)

var (
	blockService    service.IBlockService    = new(service.BlockService)
	txService       service.ITxService       = new(service.TxService)
	nftClsService   service.INFTClsService   = new(service.NFTClsService)
	nftService      service.INFTService      = new(service.NFTService)
)