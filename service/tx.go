package service

import (
	"github.com/FiiLabs/block_explorer/errors"
	"github.com/FiiLabs/block_explorer/model/vo"
)

type ITxService interface {
	QueryTransactionByHash(hash string) (*vo.TxResp, errors.Error)
	QueryTxsByHeight(height string) (*vo.TxsResp, errors.Error)
	GetTxsCount() (int64, errors.Error)
	GetTxs(p_page string, p_size string) (*vo.TxsResp, errors.Error)
	GetLTxs() (*vo.TxsResp, errors.Error)
}

var _ ITxService = new(TxService)

type TxService struct {
}

func (svc *TxService) QueryTransactionByHash(hash string) (*vo.TxResp, errors.Error) {

	tx,err := txRepo.GetTransactionByHash(hash)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return &vo.TxResp{
		TxId:     tx.TxId,
		TxTime:   tx.Time,
		Height:   tx.Height,
		TxHash:   tx.TxHash,
		Type:     tx.Type,
		Memo:     tx.Memo,
		Status:   tx.Status,
		Log:      tx.Log,
		Fee:      tx.Fee,
		GasUsed:  tx.Fee.Gas,
		Types:    tx.Types,
		EventsNew: tx.EventsNew,
		Signers:  tx.Signers,
		Msgs:     tx.DocTxMsgs,
		Addrs:    tx.Addrs,
		TxIndex:  tx.TxIndex,
	}, nil
}

func (svc *TxService) QueryTxsByHeight(height string) (*vo.TxsResp, errors.Error) {

	txBatch,err := txRepo.GetTxsByHeight(height)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	txs := make(vo.TxsResp, len(txBatch))
	for i, tx := range txBatch {
	txs[i].TxId = tx.TxId
	txs[i].TxTime = tx.Time
	txs[i].Height = tx.Height
	txs[i].TxHash = tx.TxHash
	txs[i].Type = tx.Type
	txs[i].Memo = tx.Memo
	txs[i].Status = tx.Status
	txs[i].Log = tx.Log
	txs[i].Fee = tx.Fee
	txs[i].GasUsed = tx.Fee.Gas
	txs[i].Types = tx.Types
	txs[i].EventsNew = tx.EventsNew
	txs[i].Signers = tx.Signers
	txs[i].Msgs = tx.DocTxMsgs
	txs[i].Addrs = tx.Addrs
	txs[i].TxIndex = tx.TxIndex
	}
	return &txs, nil
}

func (svc *TxService) GetTxsCount() (int64, errors.Error) {

	n,err := txRepo.GetTxsCount()
	if err != nil {
		return 0, errors.Wrap(err)
	}

	return n, nil
}
func (svc *TxService) GetTxs(p_page string, p_size string) (*vo.TxsResp, errors.Error) {

	txBatch,err := txRepo.GetTxs(p_page, p_size)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	txs := make(vo.TxsResp, len(txBatch))
	for i, tx := range txBatch {
		txs[i].TxId = tx.TxId
		txs[i].TxTime = tx.Time
		txs[i].Height = tx.Height
		txs[i].TxHash = tx.TxHash
		txs[i].Type = tx.Type
		txs[i].Memo = tx.Memo
		txs[i].Status = tx.Status
		txs[i].Log = tx.Log
		txs[i].Fee = tx.Fee
		txs[i].GasUsed = tx.Fee.Gas
		txs[i].Types = tx.Types
		txs[i].EventsNew = tx.EventsNew
		txs[i].Signers = tx.Signers
		txs[i].Msgs = tx.DocTxMsgs
		txs[i].Addrs = tx.Addrs
		txs[i].TxIndex = tx.TxIndex
	}
	return &txs, nil
}
func (svc *TxService) GetLTxs() (*vo.TxsResp, errors.Error) {

	txBatch,err := txRepo.GetLTxs()
	if err != nil {
		return nil, errors.Wrap(err)
	}
	txs := make(vo.TxsResp, len(txBatch))
	for i, tx := range txBatch {
		txs[i].TxId = tx.TxId
		txs[i].TxTime = tx.Time
		txs[i].Height = tx.Height
		txs[i].TxHash = tx.TxHash
		txs[i].Type = tx.Type
		txs[i].Memo = tx.Memo
		txs[i].Status = tx.Status
		txs[i].Log = tx.Log
		txs[i].Fee = tx.Fee
		txs[i].GasUsed = tx.Fee.Gas
		txs[i].Types = tx.Types
		txs[i].EventsNew = tx.EventsNew
		txs[i].Signers = tx.Signers
		txs[i].Msgs = tx.DocTxMsgs
		txs[i].Addrs = tx.Addrs
		txs[i].TxIndex = tx.TxIndex
	}
	return &txs, nil
}