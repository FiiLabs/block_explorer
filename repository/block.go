package repository

import (
	"context"
	"github.com/FiiLabs/block_explorer/models"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)

type IBlockRepo interface {
	GetBlockByHeight(height string) (*models.Block, error)
	GetBlockByHash(hash string) (*models.Block, error)
	GetBlocks(t_page string, t_size string) ([]*models.Block, error)
}
var _ IBlockRepo = new(BlockRepo)

type BlockRepo struct {
}

func (repo *BlockRepo) coll() *qmgo.Collection {
	return models.GetClient().Database(models.GetDbConf().Database).Collection(models.Block{}.Name())
}

func (repo *BlockRepo) GetBlockByHeight(height string) (*models.Block, error) {
	var res *models.Block

	heightT, errT := strconv.ParseInt(height, 10, 64)
	if errT != nil {
		return nil, errT
	}
	query := bson.M{
		"height": heightT,
	}
	err := repo.coll().Find(context.Background(), query).One(&res)
	return res, err
}

func (repo *BlockRepo) GetBlockByHash(hash string) (*models.Block, error) {
	var res *models.Block

	query := bson.M{
		"hash": hash,
	}
	err := repo.coll().Find(context.Background(), query).One(&res)
	return res, err
}

func (repo *BlockRepo) GetBlocks(t_page string, t_size string) ([]*models.Block, error) {
	var res []*models.Block
	pageT, errT := strconv.ParseInt(t_page, 10, 64)
	if errT != nil {
		return nil, errT
	}
	sizeT, errT := strconv.ParseInt(t_size, 10, 64)
	if errT != nil {
		return nil, errT
	}
	err := repo.coll().Find(context.Background(), bson.M{}).Sort("height").Skip(pageT*10).Limit(sizeT).All(&res)
	return res, err
}