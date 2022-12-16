package service
import "github.com/FiiLabs/block_explorer/repository"

var (
	blockRepo            repository.IBlockRepo            = new(repository.BlockRepo)
	txRepo               repository.ITxRepo               = new(repository.TxRepo)
	nftClsRepo           repository.INFTClsRepo           = new(repository.NFTClsRepo)
	nftRepo              repository.INFTRepo              = new(repository.NFTRepo)
)