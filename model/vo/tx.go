package vo

import (
	"github.com/FiiLabs/block_explorer/models"
)

type TxResp struct {
	TxId     int64             `json:"tx_id"`
	TxTime   int64             `json:"time"`
	Height   int64             `json:"height"`
	TxHash   string            `json:"tx_hash"`
	Type     string            `json:"type"`
	Memo     string            `json:"memo"`
	Status   uint32            `json:"status"`
	Log      string            `json:"log"`
	Fee      *models.Fee               `json:"fee"`
	GasUsed  int64             `json:"gas_used"`
	Types    []string          `json:"types"`
	EventsNew     []*models.EventNew        `json:"events_new"`
	Signers  []string          `json:"signers"`
	Msgs     []*models.TxMsg    `json:"msgs"`
	Addrs    []string          `json:"addrs"`
	TxIndex   uint32           `json:"tx_index"`
}
type TxsResp []TxResp
