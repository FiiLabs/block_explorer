package service
import "github.com/FiiLabs/block_explorer/repository"

var (
	blockRepo            repository.IBlockRepo            = new(repository.BlockRepo)
)