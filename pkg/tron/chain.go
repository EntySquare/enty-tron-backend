package tron

import (
	"context"
	"database/sql"
	"entysquare/enty-tron-backend/pkg/util"
	"entysquare/enty-tron-backend/storage"
	"entysquare/enty-tron-backend/storage/sqlutil"
	"strconv"
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
		println(time.Now().Format("2006-01-02::15:04:05"))
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
				} else if err2 == nil {
					txs.Status = "-1"
					err = db.UpdateTxsByHash(ctx, txn, txs)
					if err != nil {
						return err
					}
				}
			}
			addrL, err := db.ListAddressByStatus(ctx, txn)
			if err != nil {
				return err
			}
			for address, addr := range addrL {
				ut, err := strconv.ParseInt(addr.UpdateTime, 10, 64)
				if err != nil {
					return err
				}
				diff := time.Now().Unix() - ut
				if diff > 150 {
					if addr.Nft == "1" {
						txs, err := db.SelectTxsByAddressAndType(ctx, txn, address, "2")
						if err != nil {
							return err
						}
						if txs == nil {
							addr.Nft = "0"
							err = db.UpdateAddressById(ctx, txn, addr)
							if err != nil {
								return err
							}
						}
					}
					if addr.Tb == "1" {
						txs, err := db.SelectTxsByAddressAndType(ctx, txn, address, "1")
						if err != nil {
							return err
						}
						if txs == nil {
							addr.Tb = "0"
							err = db.UpdateAddressById(ctx, txn, addr)
							if err != nil {
								return err
							}
						}
					}
				}
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 20)
	}
}
