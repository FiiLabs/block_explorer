package service

import (
	"github.com/FiiLabs/block_explorer/errors"
	"github.com/FiiLabs/block_explorer/model/vo"
)
type IBlockService interface {
	Query(hash string, req vo.BlockReq) (*vo.BlockResp, errors.Error)
}

var _ IBlockService = new(BlockService)

type BlockService struct {
}

func (svc *BlockService) Query(hash string, req vo.BlockReq) (*vo.BlockResp, errors.Error) {
	return &vo.BlockResp{
		Height: 1,
		Hash:hash,
		Txn:2,
		Time:3252352,
		Proposer:"sdsdgadsg",
	}, nil
}