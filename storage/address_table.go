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
    nft       varchar(20),
    tb       varchar(20),
	update_time TEXT
);
`
const insertAddressSQL = "" +
	"INSERT INTO address ( address,nft,tb,update_time ) VALUES ( $1, $2, $3 , $4) "
const selectAddressSQL = "" +
	"SELECT id,address,nft,tb,update_time FROM address where id = $1 "
const selectAddressByAddressSQL = "" +
	"SELECT id,address,nft,tb,update_time FROM address where address = $1 "
const selectNftSoldSQL = "" +
	"SELECT count(1) FROM address where  nft = '1'"
const selectTbSoldSQL = "" +
	"SELECT count(1) FROM address where  tb = '1'"
const listAddressByStatusSQL = "" +
	"SELECT id,address,nft,tb,update_time FROM address where nft = '1' or tb = '1' "
const updateAddressByIdSQL = "" +
	" UPDATE address SET address = $1 , nft = $2 , tb = $3 , update_time = $4 " +
	" WHERE " +
	" id = $5 "

type addressSta struct {
	insertAddressStmt          *sql.Stmt
	selectAddressStmt          *sql.Stmt
	updateAddressByIdStmt      *sql.Stmt
	selectNftSoldStmt          *sql.Stmt
	selectTbSoldStmt           *sql.Stmt
	listAddressByStatusStmt    *sql.Stmt
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
	if s.selectNftSoldStmt, err = db.Prepare(selectNftSoldSQL); err != nil {
		return
	}
	if s.selectTbSoldStmt, err = db.Prepare(selectTbSoldSQL); err != nil {
		return
	}
	if s.listAddressByStatusStmt, err = db.Prepare(listAddressByStatusSQL); err != nil {
		return
	}
	if s.insertAddressStmt, err = db.Prepare(insertAddressSQL); err != nil {
		return
	}
	if s.updateAddressByIdStmt, err = db.Prepare(updateAddressByIdSQL); err != nil {
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
			&t1.Nft,
			&t1.Tb,
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
			&t1.Nft,
			&t1.Tb,
			&t1.UpdateTime)
		if t1.Id != 0 {
			return &t1, nil
		}
	}
	return nil, err
}

//user_id, address,key,nft,tb,chain_name,update_time
func (s *addressSta) insertAddress(ctx context.Context, txn *sql.Tx, addr types.Address) (err error) {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	res, err := sqlutil.TxStmt(txn, s.insertAddressStmt).
		ExecContext(ctx, addr.Address, addr.Nft, addr.Tb, now)
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
func (s *addressSta) updateAddressById(ctx context.Context, txn *sql.Tx, b types.Address) (err error) {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	res, err := sqlutil.TxStmt(txn, s.updateAddressByIdStmt).Exec(
		b.Address, b.Nft, b.Tb, now, //set
		b.Id, //where
	)
	if err != nil {
		return err
	}
	i, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}
	if i == 0 {
		return fmt.Errorf("updateAddressById insertBalance i==0")
	}
	return nil
}
func (s *addressSta) selectAllSold(
	ctx context.Context, txn *sql.Tx,
) (int64, int64, error) {
	rows, err := sqlutil.TxStmt(txn, s.selectTbSoldStmt).QueryContext(ctx)
	defer rows.Close()
	if err != nil {
		return 0, 0, err
	}
	var b int64
	for rows.Next() {
		if err = rows.Scan(
			&b,
		); err != nil {
			return 0, 0, err
		}
	}
	rows, err = sqlutil.TxStmt(txn, s.selectNftSoldStmt).QueryContext(ctx)
	defer rows.Close()
	if err != nil {
		return 0, 0, err
	}
	//id,user_id,address,key,types,update_time
	var n int64
	for rows.Next() {
		if err = rows.Scan(
			&n,
		); err != nil {
			return 0, 0, err
		}
	}
	return b, n, rows.Err()
}
func (s *addressSta) listAddressByStatus(
	ctx context.Context, txn *sql.Tx,
) ([]types.Address, error) {
	rows, err := sqlutil.TxStmt(txn, s.listAddressByStatusStmt).QueryContext(ctx)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	maps := make([]types.Address, 0)
	for rows.Next() {
		var b types.Address
		if err = rows.Scan(
			&b.Id,
			&b.Address,
			&b.Nft,
			&b.Tb,
			&b.UpdateTime,
		); err != nil {
			return nil, err
		}
		maps = append(maps, b)
	}
	return maps, rows.Err()
}
