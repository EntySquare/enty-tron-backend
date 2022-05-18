package types

type Address struct {
	Id         int64
	Address    string
	Types      string // 0 新账户 1 已购买1类产品账户  2 已购买2类产品账户
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
