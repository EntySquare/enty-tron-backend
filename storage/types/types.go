package types

type Address struct {
	Id         int64
	Address    string
	Types      string //
	UpdateTime string
}
type Txs struct {
	Id         int64
	Hash       *string
	Chain      string
	Status     string
	InsertTime string
}
type User struct {
	Id     int64
	Phone  string
	Cypher string
}
