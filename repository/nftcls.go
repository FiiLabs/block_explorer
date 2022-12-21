package repository

import (
	"context"
	"github.com/FiiLabs/block_explorer/models"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)

type INFTClsRepo interface {
	GetNFTCls(t_page string, t_size string) ([]*models.NFTCls, error)
	GetNFTClsCount() (int64, error)
}
var _ INFTClsRepo = new(NFTClsRepo)

type NFTClsRepo struct {
}

func (repo *NFTClsRepo) coll() *qmgo.Collection {
	return models.GetClient().Database(models.GetDbConf().Database).Collection(models.NFTCls{}.Name())
}
func (repo *NFTClsRepo) GetNFTCls(t_page string, t_size string) ([]*models.NFTCls, error) {
	var res []*models.NFTCls
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

func (repo *NFTClsRepo) GetNFTClsCount() (int64, error) {

	n,err := repo.coll().Find(context.Background(), bson.M{}).Count()
	return n, err
}