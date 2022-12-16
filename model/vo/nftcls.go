package vo

type NFTClsResp struct {
	Id                string  `json:"id"`
	ClsName           string  `json:"name"`
	Schema            string  `json:"schema"`
	Sender            string  `json:"sender"`
	Symbol            string  `json:"symbol"`
	MintRestricted    bool    `json:"mintRestricted"`
	UpdateRestricted  bool    `json:"updateRestricted"`
	Description       string  `json:"description"`
	Uri               string  `json:"uri"`
	UriHash           string  `json:"uriHash"`
	Data              string  `json:"data"`
	Time     	      int64   `json:"time"`
}

type NFTClsesResp []NFTClsResp
