package routing

type CheckAddressReq struct {
	Address string `json:"address"`
}
type CheckAddressResp struct {
	RetCode  string `json:"retCode"`
	TbLimit  string `json:"tbLimit"`  // 0 未参与 1 已锁定待支付 2 已支付待确认 3已确认
	NftLimit string `json:"nftLimit"` // 0 未参与 1 已锁定待支付 2 已支付待确认 3已确认
}
type ListAddressJoinedReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type ListAddressJoinedResp struct {
	RetCode string          `json:"retCode"`
	List    []AddressStatus `json:"list"`
}
type SetAddressStatusReq struct {
	Username   string        `json:"username"`
	Password   string        `json:"password"`
	AddrStatus AddressStatus `json:"addrStatus"`
}
type SetAddressStatusResp struct {
	RetCode string `json:"retCode"`
	Message string `json:"message"`
}
type AddressStatus struct {
	Address         string `json:"address"`
	TransactionId   string `json:"transactionId"`
	TransactionType string `json:"transactionType"`
	Status          string `json:"status"` // 0 已支付 1 链上已确认 2 管理员已确认

}
type TransactionInsertReq struct {
	TransactionId   string `json:"transactionId"`
	Address         string `json:"address"`
	TransactionType string `json:"transactionType"` //交易类型，tb coin 1，nft 2
}
type ConfirmLimitReq struct {
	Address         string `json:"address"`
	TransactionType string `json:"transactionType"` //交易类型，tb coin 1，nft 2
}
type ConfirmLimitResp struct {
	RetCode string `json:"retCode"`
	Message string `json:"message"`
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
