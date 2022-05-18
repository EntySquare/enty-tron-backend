package storage

import (
	"context"
	"database/sql"
	"entysquare/enty-tron-backend/storage/sqlutil"
	"entysquare/enty-tron-backend/storage/types"
	"fmt"
	"strconv"
	"time"
)

const addressSchema = `
create table IF NOT EXISTS address
(
    id  serial PRIMARY KEY,
    address   varchar(255),
    types       varchar(20),
	update_time TEXT
);
`
const insertAddressSQL = "" +
	"INSERT INTO address ( address,types,update_time ) VALUES ( $1, $2, $3 ) "
const selectAddressSQL = "" +
	"SELECT id,address,types,update_time FROM address where id = $1 "
const selectAddressByAddressSQL = "" +
	"SELECT id,address,types,update_time FROM address where address = $1 "

type addressSta struct {
	insertAddressStmt          *sql.Stmt
	selectAddressStmt          *sql.Stmt
	selectAddressByAddressStmt *sql.Stmt
}

func (s *addressSta) execSchema(db *sql.DB) error {
	_, err := db.Exec(addressSchema)
	return err
}

func (s *addressSta) prepare(db *sql.DB) (err error) {
	if s.selectAddressStmt, err = db.Prepare(selectAddressSQL); err != nil {
		return
	}
	if s.selectAddressByAddressStmt, err = db.Prepare(selectAddressByAddressSQL); err != nil {
		return
	}
	if s.insertAddressStmt, err = db.Prepare(insertAddressSQL); err != nil {
		return
	}
	return
}
func (s *addressSta) selectAddress(
	ctx context.Context, txn *sql.Tx, id int64,
) (*types.Address, error) {
	rows, err := sqlutil.TxStmt(txn, s.selectAddressStmt).QueryContext(ctx, id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var t1 types.Address
	if rows.Next() {
		err = rows.Scan(
			&t1.Id,
			&t1.Address,
			&t1.Types,
			&t1.UpdateTime)
		if t1.Id != 0 {
			return &t1, nil
		}
	}
	return nil, err
}
func (s *addressSta) selectAddressByAddress(
	ctx context.Context, txn *sql.Tx, address string,
) (*types.Address, error) {
	rows, err := sqlutil.TxStmt(txn, s.selectAddressByAddressStmt).QueryContext(ctx, address)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var t1 types.Address
	if rows.Next() {
		err = rows.Scan(
			&t1.Id,
			&t1.Address,
			&t1.Types,
			&t1.UpdateTime)
		if t1.Id != 0 {
			return &t1, nil
		}
	}
	return nil, err
}

//user_id, address,key,types,chain_name,update_time
func (s *addressSta) insertAddress(ctx context.Context, txn *sql.Tx, addr types.Address) (err error) {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	res, err := sqlutil.TxStmt(txn, s.insertAddressStmt).
		ExecContext(ctx, addr.Address, addr.Types, now)
	if err != nil {
		return err
	}
	i, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if i == 0 {
		return fmt.Errorf("insertUser insertAddress i==0")
	}
	return nil
}
