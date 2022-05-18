package storage

import (
	"context"
	"database/sql"
	"entysquare/enty-tron-backend/conf"
	"entysquare/enty-tron-backend/storage/types"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Database struct {
	Db  *sql.DB
	add addressSta
	txs txsSta
}

// NewDatabase creates a new accounts and profiles database
func NewDatabase() (*Database, error) {
	var result Database
	var err error
	conf := new(conf.Conf).GetConf()
	fmt.Println(conf)
	fmt.Println("CONF:", conf)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Postgres.Host, conf.Postgres.Port, conf.Postgres.User, conf.Postgres.Password, conf.Postgres.Dbname)

	fmt.Println("psqlInfo:", psqlInfo)
	result.Db, err = sql.Open("pgx", psqlInfo)

	if err != nil {
		panic(err)
	}
	//result.writer = sqlutil.NewWriterDb()

	//是否创建表
	if conf.Postgres.ExecSchema {
		if err = result.add.execSchema(result.Db); err != nil {
			fmt.Println(err)
			return nil, err
		}
		if err = result.txs.execSchema(result.Db); err != nil {
			fmt.Println(err)
			return nil, err
		}

	}

	if err = result.add.prepare(result.Db); err != nil {
		return nil, err
	}
	if err = result.txs.prepare(result.Db); err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *Database) SelectAddress(ctx context.Context, txn *sql.Tx, id int64,
) (*types.Address, error) {
	return d.add.selectAddress(ctx, txn, id)
}
func (d *Database) SelectAddressByAddress(ctx context.Context, txn *sql.Tx, address string,
) (*types.Address, error) {
	return d.add.selectAddressByAddress(ctx, txn, address)
}
func (d *Database) InsertAddress(ctx context.Context, txn *sql.Tx, addr types.Address,
) error {
	return d.add.insertAddress(ctx, txn, addr)
}
func (d *Database) UpdateAddressById(ctx context.Context, txn *sql.Tx, a types.Address) (err error) {
	return d.add.updateAddressById(ctx, txn, a)
}
func (d *Database) InsertTxs(ctx context.Context, txn *sql.Tx, tx types.Txs) (err error) {
	return d.txs.insertTxs(ctx, txn, tx)
}
func (d *Database) UpdateTxs(ctx context.Context, txn *sql.Tx, b types.Txs) (err error) {
	return d.txs.updateTxs(ctx, txn, b)
}
func (d *Database) SelectTxs(ctx context.Context, txn *sql.Tx, hash string, chain string) (result *types.Txs, err error) {
	return d.txs.selectTxs(ctx, txn, hash)
}
func (d *Database) SelectCollectTxs(ctx context.Context, txn *sql.Tx, chain string) (result int64, err error) {
	return d.txs.selectCollectTxs(ctx, txn)
}
func (d *Database) UpdateTxsByHash(ctx context.Context, txn *sql.Tx, b types.Txs) (err error) {
	return d.txs.updateTxsByHash(ctx, txn, b)
}
