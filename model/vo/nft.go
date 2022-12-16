package vo

type NFTResp struct {
	Id                string  `json:"id"`
	NTFName           string  `json:"name"`
	DenomId           string  `json:"denomId"`
	URI               string  `json:"uri"`
	Data              string  `json:"data"`
	Sender            string  `json:"sender"`
	Recipient         string  `json:"recipient"`
	UriHash           string  `json:"uriHash"`
	Time     	      int64   `json:"time"`
}

type NFTsResp []NFTResp
