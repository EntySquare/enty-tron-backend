package tron

import (
	"context"
	"database/sql"
	"entysquare/enty-tron-backend/pkg/util"
	"entysquare/enty-tron-backend/storage"
	"entysquare/enty-tron-backend/storage/sqlutil"
	"time"
)

//type TH struct {
//	ctx          context.Context
//	StreamClient *stream.Client
//	db           *storage.Database
//}
//
//func Build(db *storage.Database) (*TH, error) {
//	cli, err := stream.GenClient()
//	if err != nil {
//		panic("gen kafka client error")
//	}
//	c := &TH{
//		ctx:          context.TODO(),
//		StreamClient: cli,
//		db:           db,
//	}
//	//	err = c.StreamClient.Register(stream.WITHDRAW_COLLECT, c.WithdrawCollectHandler)
//	//err = c.StreamClient.Register(stream.TRANSACTION, c.TransactionHandler)
//	//if err != nil {
//	//	return nil, err
//	//}
//	return c, nil
//}
//func (c *TH) Run() error {
//	//for {
//	//	err := r.RiskHandler(conf, cli)
//	//	if err != nil {
//	//		return err
//	//	}
//	return c.StreamClient.Process()
//	//}
//}
func ScanTron(db *storage.Database) {
	for {
		//println(time.Now().Format("2006-01-02::15:04:05"))
		err := sqlutil.WithTransaction(db.Db, func(txn *sql.Tx) error {
			ctx := context.TODO()
			//start := time.Now()
			hashL, err := db.ListTxsByStatus(ctx, txn)
			if err != nil {
				return err
			}
			for hash, txs := range hashL {
				flag, err2 := util.CheckTransaction(hash)
				if flag {
					txs.Status = "1"
					err = db.UpdateTxsByHash(ctx, txn, txs)
					if err != nil {
						return err
					}
				}
				if !flag && err2 == nil {
					txs.Status = "-1"
					err = db.UpdateTxsByHash(ctx, txn, txs)
					if err != nil {
						return err
					}
					addr, err := db.SelectAddressByAddress(ctx, txn, txs.Address)
					if err != nil {
						return err
					}
					if txs.TransactionType == "1" {
						addr.Tb = "0"
					}
					if txs.TransactionType == "2" {
						addr.Nft = "0"
					}
					err = db.UpdateAddressById(ctx, txn, *addr)
					if err != nil {
						return err
					}
				}
				if !flag && err2 != nil {
					return err2
				}
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 2)
	}
}
