package service

import (
	"github.com/FiiLabs/block_explorer/errors"
	"github.com/FiiLabs/block_explorer/model/vo"
	sdk "github.com/cosmos/cosmos-sdk/types"
)
type IBlockService interface {
	QueryBlockByHeight(height string) (*vo.BlockResp, errors.Error)
	QueryBlockByHash(hash string) (*vo.BlockResp, errors.Error)
	QueryBlocks(p_num string) (*vo.BlocksResp, errors.Error)
	QueryLBlocks() (*vo.BlocksResp, errors.Error)
}

var _ IBlockService = new(BlockService)

type BlockService struct {
}

func (svc *BlockService) QueryBlockByHeight(height string) (*vo.BlockResp, errors.Error) {

	block,err := blockRepo.GetBlockByHeight(height)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	providerAddr, err := sdk.ValAddressFromHex(block.Proposer)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return &vo.BlockResp{
		Height: block.Height,
		Hash:block.Hash,
		Txn:block.Txn,
		Time:block.Time,
		Proposer:providerAddr.String(),
	}, nil
}

func (svc *BlockService) QueryBlockByHash(hash string) (*vo.BlockResp, errors.Error) {

	block,err := blockRepo.GetBlockByHash(hash)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	providerAddr, err := sdk.ValAddressFromHex(block.Proposer)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return &vo.BlockResp{
		Height: block.Height,
		Hash:block.Hash,
		Txn:block.Txn,
		Time:block.Time,
		Proposer:providerAddr.String(),
	}, nil
}

func (svc *BlockService) QueryBlocks(p_num string) (*vo.BlocksResp, errors.Error) {

	blocks,err := blockRepo.GetBlocks(p_num)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	blks := make(vo.BlocksResp, len(blocks))
	for i, v := range blocks {
		blks[i].Txn = v.Txn
		blks[i].Height = v.Height
		blks[i].Hash = v.Hash
		blks[i].Time = v.Time
		providerAddr, err := sdk.ValAddressFromHex(v.Proposer)
		if err != nil {
			return nil, errors.Wrap(err)
		}
		blks[i].Proposer = providerAddr.String()
	}
	return &blks, nil
}
func (svc *BlockService) QueryLBlocks() (*vo.BlocksResp, errors.Error) {

	blocks,err := blockRepo.GetLBlocks()
	if err != nil {
		return nil, errors.Wrap(err)
	}
	blks := make(vo.BlocksResp, len(blocks))
	for i, v := range blocks {
		blks[i].Txn = v.Txn
		blks[i].Height = v.Height
		blks[i].Hash = v.Hash
		blks[i].Time = v.Time
		providerAddr, err := sdk.ValAddressFromHex(v.Proposer)
		if err != nil {
			return nil, errors.Wrap(err)
		}
		blks[i].Proposer = providerAddr.String()
	}
	return &blks, nil
}