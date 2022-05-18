package routing

type CheckAddressReq struct {
	Address         string `json:"address"`
	TransactionType string `json:"transactionType"` //交易类型，tb coin 1，nft 2
}
type CheckAddressResp struct {
	RetCode string `json:"retCode"`
	Message string `json:"message"`
}
type TransactionInsertReq struct {
	TransactionId   string `json:"transactionId"`
	Address         string `json:"address"`
	TransactionType string `json:"transactionType"` //交易类型，tb coin 1，nft 2
}
type TransactionInsertResp struct {
	RetCode string `json:"retCode"`
	Message string `json:"message"`
}
type QueryCoinLimitResp struct {
	RetCode        string `json:"retCode"`
	TbHasBeenSold  int64  `json:"tbHasBeenSold"`
	NftHasBeenSold int64  `json:"nftHasBeenSold"`
}
