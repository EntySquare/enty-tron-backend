package tron

import (
	"context"
	"database/sql"
	"entysquare/enty-tron-backend/pkg/util"
	"entysquare/enty-tron-backend/storage"
	"entysquare/enty-tron-backend/storage/sqlutil"
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
func ScanTron(db *storage.Database, address string, ttype string) {
	//println(time.Now().Format("2006-01-02::15:04:05"))
	//time.Sleep(time.Second * 20)
	ctx := context.TODO()
	err := sqlutil.WithTransaction(db.Db, func(txn *sql.Tx) error {
		//start := time.Now()
		txs, err := db.SelectTxsByAddressAndType(ctx, txn, address, ttype)
		if err != nil {
			return err
		}
		hash := *txs.Hash
		flag, _ := util.CheckTransaction(hash)
		if flag {
			txs.Status = "1"
			err = db.UpdateTxsByHash(ctx, txn, *txs)
			if err != nil {
				return err
			}
		}
		if !flag {
			txs.Status = "-1"
			err = db.UpdateTxsByHash(ctx, txn, *txs)
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

		return nil
	})
	if err != nil {
		panic(err)
	}

}
