// interface for a document

package models

var (
	SyncTaskModel SyncTask
	BlockModel    Block
	TxModel       Tx
	NFTClsModel   NFTCls
	NFTModel      NFT

	Collections = []Docs{
		SyncTaskModel,
		BlockModel,
		TxModel,
		NFTClsModel,
		NFTModel,
	}
)

type (
	Docs interface {
		// collection name
		Name() string
		// ensure indexes
		EnsureIndexes()
		// primary key pair(used to find a unique record)
		PkKvPair() map[string]interface{}
	}
)

func ensureDocsIndexes() {
	if len(Collections) > 0 {
		for _, v := range Collections {
			v.EnsureIndexes()
		}
	}
}
