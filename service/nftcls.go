package service

import (
	"github.com/FiiLabs/block_explorer/errors"
	"github.com/FiiLabs/block_explorer/model/vo"
)

type INFTClsService interface {
	QueryNFTCls(p_page string, p_size string) (*vo.NFTClsesResp, errors.Error)
}

var _ INFTClsService = new(NFTClsService)

type NFTClsService struct {
}

func (svc *NFTClsService) QueryNFTCls(p_page string, p_size string) (*vo.NFTClsesResp, errors.Error) {

	nftBatch,err := nftClsRepo.GetNFTCls(p_page, p_size)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	nftclses := make(vo.NFTClsesResp, len(nftBatch))
	for i, v := range nftBatch {
		nftclses[i].Id = v.Id
		nftclses[i].ClsName = v.ClsName
		nftclses[i].Schema = v.Schema
		nftclses[i].Sender = v.Sender
		nftclses[i].Symbol = v.Symbol
		nftclses[i].MintRestricted = v.MintRestricted
		nftclses[i].UpdateRestricted = v.UpdateRestricted
		nftclses[i].Description = v.Description
		nftclses[i].Uri = v.Uri
		nftclses[i].UriHash = v.UriHash
		nftclses[i].Data = v.Data
		nftclses[i].Time = v.Time
	}
	return &nftclses, nil
}