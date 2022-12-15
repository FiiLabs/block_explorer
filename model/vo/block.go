package vo

type BlockReq struct {
	Height          string `json:"height"`
}

type BlockResp struct {
	Height   int64  `json:"height"`
	Hash     string `json:"hash"`
	Txn      int64  `json:"txn"`
	Time     int64  `json:"time"`
	Proposer string `json:"proposer"`
}