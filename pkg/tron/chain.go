package tron

import (
	"context"
	"database/sql"
	"entysquare/enty-tron-backend/pkg/util"
	"entysquare/enty-tron-backend/storage"
	"entysquare/enty-tron-backend/storage/sqlutil"
	"fmt"
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
		for i := 0; i < 2; i++ {
			time.Sleep(time.Second * 10)
			flag, _ := util.CheckTransaction(hash)
			if flag {
				txs.Status = "1"
				err = db.UpdateTxsByHash(ctx, txn, *txs)
				if err != nil {
					return err
				}
				return nil
			}
		}
		flag, _ := util.CheckTransaction(hash)
		if flag {
			txs.Status = "1"
			err = db.UpdateTxsByHash(ctx, txn, *txs)
			if err != nil {
				return err
			}
			return nil
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
		fmt.Println(err)
	}
	fmt.Println(address + " ::::updateTransaction:::: ")
}
func UnlockLimit(db *storage.Database, address string) {
	time.Sleep(time.Second * 60)
	ctx := context.TODO()
	err := sqlutil.WithTransaction(db.Db, func(txn *sql.Tx) error {
		addr, err := db.SelectAddressByAddress(ctx, txn, address)
		if err != nil {
			return err
		}
		var count = 0
		if addr.Nft == "1" {
			txs, err := db.SelectTxsByAddressAndType(ctx, txn, addr.Address, "2")
			if err != nil {
				return err
			}
			if txs == nil {
				count += 1
				fmt.Println(addr.Address + " ::::return limit:::: 2")
			}
		}
		if addr.Tb == "1" {
			txs, err := db.SelectTxsByAddressAndType(ctx, txn, addr.Address, "1")
			if err != nil {
				return err
			}
			if txs == nil {
				count += 2

				fmt.Println(addr.Address + " ::::return limit:::: 1")
			}
		}
		if count == 3 {
			addr.Tb = "0"
			addr.Nft = "0"
		} else if count == 2 {
			addr.Tb = "0"
		} else if count == 1 {
			addr.Nft = "0"
		}
		err = db.UpdateAddressById(ctx, txn, *addr)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
