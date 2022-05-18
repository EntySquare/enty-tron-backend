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

const txsSchema = `
create table IF NOT EXISTS txs
(
    id  serial PRIMARY KEY,
    hash        text,
    -- success,fail
    status      	 text,
    chain_name       text,
    insert_time TEXT
);
`

type txsSta struct {
	insertTxsyStmt       *sql.Stmt
	updateTxsStmt        *sql.Stmt
	updateTxsByHashStmt  *sql.Stmt
	selectTxsStmt        *sql.Stmt
	selectCollectTxsStmt *sql.Stmt
}

const insertTxsySQL = "" +
	"INSERT INTO txs(hash, status, chain_name, insert_time) VALUES ($1, $2, $3, $4) "
const updateTxsSQL = "" +
	" UPDATE txs SET hash = $1 , status = $2 , chain_name = $3 , insert_time = $4" +
	" WHERE " +
	" id = $5  "
const selectTxsSQL = "" +
	"SELECT id,hash,status,chain_name,insert_time FROM txs where chain_name = $1 and hash =$2 "
const selectCollectTxsSQL = "" +
	"SELECT count(1) FROM txs where chain_name = $1 and status = 'collect'"
const updateTxsByHashSQL = "" +
	" UPDATE txs SET hash = $1 , status = $2 , chain_name = $3 , insert_time = $4" +
	" WHERE " +
	" hash = $5  "

func (s *txsSta) execSchema(db *sql.DB) error {
	_, err := db.Exec(txsSchema)
	return err
}

func (s *txsSta) prepare(db *sql.DB) (err error) {
	if s.insertTxsyStmt, err = db.Prepare(insertTxsySQL); err != nil {
		return
	}
	if s.updateTxsStmt, err = db.Prepare(updateTxsSQL); err != nil {
		return
	}
	if s.updateTxsByHashStmt, err = db.Prepare(updateTxsByHashSQL); err != nil {
		return
	}
	if s.selectTxsStmt, err = db.Prepare(selectTxsSQL); err != nil {
		return
	}
	if s.selectCollectTxsStmt, err = db.Prepare(selectCollectTxsSQL); err != nil {
		return
	}
	return
}

//hash, form_, to_, num, chain, contract_id, insert_time
func (s *txsSta) insertTxs(ctx context.Context, txn *sql.Tx, p types.Txs) (err error) {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	_, err = sqlutil.TxStmt(txn, s.insertTxsyStmt).
		ExecContext(ctx,
			p.Hash, p.Status, p.Chain, &now)
	fmt.Println(p.Status+"~txs 插入：", txn)
	return
}
func (s *txsSta) updateTxs(ctx context.Context, txn *sql.Tx, b types.Txs) (err error) {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	res, err := sqlutil.TxStmt(txn, s.updateTxsStmt).Exec(
		b.Hash, b.Status, b.Chain, now, //set
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
		return fmt.Errorf("updateTxs insertBalance i==0")
	}
	return nil
}
func (s *txsSta) selectTxs(
	ctx context.Context, txn *sql.Tx, hash string, chain string,
) (*types.Txs, error) {
	rows, err := sqlutil.TxStmt(txn, s.selectTxsStmt).QueryContext(ctx, chain, hash)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	//id,user_id,address,key,types,update_time
	var b *types.Txs
	for rows.Next() {
		if b == nil {
			b = &types.Txs{}
		}
		if err = rows.Scan(
			&b.Id,
			&b.Hash,
			&b.Status,
			&b.Chain,
			&b.InsertTime,
		); err != nil {
			return nil, err
		}
	}
	return b, rows.Err()
}
func (s *txsSta) selectCollectTxs(
	ctx context.Context, txn *sql.Tx, chain string,
) (int64, error) {
	rows, err := sqlutil.TxStmt(txn, s.selectTxsStmt).QueryContext(ctx, chain)
	defer rows.Close()
	if err != nil {
		return 0, err
	}
	//id,user_id,address,key,types,update_time
	var b int64
	for rows.Next() {
		if err = rows.Scan(
			b,
		); err != nil {
			return 0, err
		}
	}
	return b, rows.Err()
}
func (s *txsSta) updateTxsByHash(ctx context.Context, txn *sql.Tx, b types.Txs) (err error) {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	res, err := sqlutil.TxStmt(txn, s.updateTxsByHashStmt).Exec(
		b.Hash, b.Status, b.Chain, now, //set
		b.Hash, //where
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
		return fmt.Errorf("updateTxs insertBalance i==0")
	}
	return nil
}
