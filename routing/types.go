package routing

type CheckAddressResp struct {
	RetCode string `json:"retCode"`
	Message string `json:"message"`
}
type TransactionInsertReq struct {
	TransactionId   string `json:"hash"`
	Address         string `json:"address"`
	TransactionType string `json:"transactionType"` //交易类型，tb coin 1，nft 2
}
type TransactionInsertResp struct {
	RetCode string `json:"retCode"`
	Message string `json:"message"`
}
