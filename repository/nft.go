package repository

import (
	"context"
	"github.com/FiiLabs/block_explorer/models"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)

type INFTRepo interface {
	GetNFT(t_page string, t_size string) ([]*models.NFT, error)
}
var _ INFTRepo = new(NFTRepo)

type NFTRepo struct {
}

func (repo *NFTRepo) coll() *qmgo.Collection {
	return models.GetClient().Database(models.GetDbConf().Database).Collection(models.NFT{}.Name())
}
func (repo *NFTRepo) GetNFT(t_page string, t_size string) ([]*models.NFT, error) {
	var res []*models.NFT
	pageT, errT := strconv.ParseInt(t_page, 10, 64)
	if errT != nil {
		return nil, errT
	}
	sizeT, errT := strconv.ParseInt(t_size, 10, 64)
	if errT != nil {
		return nil, errT
	}
	err := repo.coll().Find(context.Background(), bson.M{}).Sort("time").Skip(pageT*10).Limit(sizeT).All(&res)
	return res, err
}
