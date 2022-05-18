package types

type Address struct {
	Id         int64
	Address    string
	Nft        string // 0 未购买 1 已购买
	Tb         string // 0 未购买 1 已购买
	UpdateTime string
}
type Txs struct {
	Id         int64
	Hash       *string
	Address    string
	Status     string
	InsertTime string
}
type User struct {
	Id     int64
	Phone  string
	Cypher string
}
