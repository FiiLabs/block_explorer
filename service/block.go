package service

import (
	"github.com/FiiLabs/block_explorer/errors"
	"github.com/FiiLabs/block_explorer/model/vo"
)
type IBlockService interface {
	QueryBlockByHeight(height string) (*vo.BlockResp, errors.Error)
	QueryBlockByHash(hash string) (*vo.BlockResp, errors.Error)
	QueryBlocks(p_page string, p_size string) (*vo.BlocksResp, errors.Error)
	QueryLBlocks() (*vo.BlocksResp, errors.Error)
}

var _ IBlockService = new(BlockService)

type BlockService struct {
}

// @Summary QueryBlockByHeight
// @Param block_height path string true "height"
// @Schemes
// @Description QueryBlockByHeight
// @Tags explorer api service
// @Accept json
// @Produce json
// @Success 200 {object} vo.BlockResp
// @Router /block/height/{block_height} [get]
func (svc *BlockService) QueryBlockByHeight(height string) (*vo.BlockResp, errors.Error) {

	block,err := blockRepo.GetBlockByHeight(height)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return &vo.BlockResp{
		Height: block.Height,
		Hash:block.Hash,
		Txn:block.Txn,
		Time:block.Time,
		Proposer:block.Proposer,
	}, nil
}

func (svc *BlockService) QueryBlockByHash(hash string) (*vo.BlockResp, errors.Error) {

	block,err := blockRepo.GetBlockByHash(hash)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return &vo.BlockResp{
		Height: block.Height,
		Hash:block.Hash,
		Txn:block.Txn,
		Time:block.Time,
		Proposer:block.Proposer,
	}, nil
}

func (svc *BlockService) QueryBlocks(p_page string, p_size string) (*vo.BlocksResp, errors.Error) {

	blocks,err := blockRepo.GetBlocks(p_page, p_size)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	blks := make(vo.BlocksResp, len(blocks))
	for i, v := range blocks {
		blks[i].Txn = v.Txn
		blks[i].Height = v.Height
		blks[i].Hash = v.Hash
		blks[i].Time = v.Time
		blks[i].Proposer = v.Proposer
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
		blks[i].Proposer = v.Proposer
	}
	return &blks, nil
}