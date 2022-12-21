package repository

import (
	"context"
	"github.com/FiiLabs/block_explorer/models"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)

type ITxRepo interface {
	GetTransactionByHash(hash string) (*models.Tx2, error)
	GetTxsByHeight(height string) ([]*models.Tx2, error)
	GetTxsCount() (int64, error)
	GetTxs(t_page string, t_size string) ([]*models.Tx2, error)
	GetLTxs() ([]*models.Tx2, error)
}
var _ ITxRepo = new(TxRepo)

type TxRepo struct {
}

func (repo *TxRepo) coll() *qmgo.Collection {
	return models.GetClient().Database(models.GetDbConf().Database).Collection(models.Tx{}.Name())
}
func (repo *TxRepo) GetTransactionByHash(hash string) (*models.Tx2, error) {
	var res models.Tx2

	query := bson.M{
		"tx_hash": hash,
	}
	err := repo.coll().Find(context.Background(), query).One(&res)
	return &res, err
}

func (repo *TxRepo) GetTxsByHeight(height string) ([]*models.Tx2, error) {
	var res []*models.Tx2
	heightT, errT := strconv.ParseInt(height, 10, 64)
	if errT != nil {
		return nil, errT
	}
	query := bson.M{
		"height": heightT,
	}
	err := repo.coll().Find(context.Background(), query).Sort("tx_index").All(&res)
	return res, err
}

func (repo *TxRepo) GetTxsCount() (int64, error) {
	n,err := repo.coll().Find(context.Background(), bson.M{}).Count()
	return n, err
}
func (repo *TxRepo) GetTxs(t_page string, t_size string) ([]*models.Tx2, error) {
	var res []*models.Tx2
	pageT, errT := strconv.ParseInt(t_page, 10, 64)
	if errT != nil {
		return nil, errT
	}
	sizeT, errT := strconv.ParseInt(t_size, 10, 64)
	if errT != nil {
		return nil, errT
	}
	err := repo.coll().Find(context.Background(), bson.M{}).Sort("-tx_id").Skip(pageT*10).Limit(sizeT).All(&res)
	return res, err
}
func (repo *TxRepo) GetLTxs() ([]*models.Tx2, error) {
	var res []*models.Tx2

	err := repo.coll().Find(context.Background(), bson.M{}).Sort("-tx_id").Limit(10).All(&res)
	return res, err
}