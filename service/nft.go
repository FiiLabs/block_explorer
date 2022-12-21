package service

import (
	"github.com/FiiLabs/block_explorer/errors"
	"github.com/FiiLabs/block_explorer/model/vo"
)

type INFTService interface {
	QueryNFT(p_page string, p_size string) (*vo.NFTsResp, errors.Error)
	GetNFTCount() (int64, errors.Error)
}

var _ INFTService = new(NFTService)

type NFTService struct {
}

func (svc *NFTService) QueryNFT(p_page string, p_size string) (*vo.NFTsResp, errors.Error) {
	nftBatch,err := nftRepo.GetNFT(p_page, p_size)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	nfts := make(vo.NFTsResp, len(nftBatch))
	for i, v := range nftBatch {
		nfts[i].Id = v.Id
		nfts[i].NTFName = v.NTFName
		nfts[i].DenomId = v.DenomId
		nfts[i].URI = v.URI
		nfts[i].Data = v.Data
		nfts[i].Sender = v.Sender
		nfts[i].Recipient = v.Recipient
		nfts[i].UriHash = v.UriHash
		nfts[i].Time = v.Time
	}
	return &nfts, nil
}

func (svc *NFTService) GetNFTCount() (int64, errors.Error) {
	n,err := nftRepo.GetNFTCount()
	if err != nil {
		return 0, errors.Wrap(err)
	}
	return n, nil
}