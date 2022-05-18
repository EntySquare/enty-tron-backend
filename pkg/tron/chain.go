package tron

import (
	"context"
	"entysquare/enty-tron-backend/conf"
	"entysquare/enty-tron-backend/storage"
)

type CH struct {
	ctx context.Context
	db  *storage.Database
}

func Build(db *storage.Database) (*CH, error) {
	c := &CH{
		ctx: context.TODO(),
		db:  db,
	}
	//	err = c.StreamClient.Register(stream.WITHDRAW_COLLECT, c.WithdrawCollectHandler)
	//err = c.StreamClient.Register(stream.TRANSACTION, c.TransactionHandler)
	return c, nil
}

func (c *CH) Run() error {
	//for {
	//	err := r.RiskHandler(conf, cli)
	//	if err != nil {
	//		return err
	//	}
	return nil
	//}
}
func ScanTron(conf *conf.Conf, db *storage.Database) {

}
