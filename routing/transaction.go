package routing

import (
	"database/sql"
	"encoding/json"
	"entysquare/enty-tron-backend/pkg/jsonerror"
	"entysquare/enty-tron-backend/pkg/util"
	"entysquare/enty-tron-backend/storage"
	"entysquare/enty-tron-backend/storage/sqlutil"
	"entysquare/enty-tron-backend/storage/types"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

var resourceLock = func() sync.Mutex {
	return sync.Mutex{}
}()

const tbLimit = 200
const nftLimit = 200

func insertTransaction(
	req *http.Request, db *storage.Database,
) util.JSONResponse {
	resourceLock.Lock()
	defer resourceLock.Unlock()
	bodyIo := req.Body
	ctx := req.Context()
	reqBody, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.NotFound("io can not been read successfully"),
		}
	}
	reqParams := &TransactionInsertReq{}
	err = json.Unmarshal(reqBody, reqParams)
	if err != nil {
		println(err)
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.Unknown("Transaction unmarshal error"),
		}
	}
	address := reqParams.Address
	err = sqlutil.WithTransaction(db.Db, func(txn *sql.Tx) error {
		tb, nft, err := db.SelectAllSold(ctx, txn)
		if err != nil {
			return err
		}
		if reqParams.TransactionType == "1" && tb == tbLimit {
			return fmt.Errorf("over limit")
		}
		if reqParams.TransactionType == "2" && nft == nftLimit {
			return fmt.Errorf("over limit")
		}
		addr, err := db.SelectAddressByAddress(ctx, nil, address)
		if err != nil {
			return err
		}
		if addr == nil {
			return fmt.Errorf("address not exist")
		} else {
			txs, err := db.SelectTxs(ctx, txn, reqParams.TransactionId)
			if txs != nil {
				return fmt.Errorf("has been insert")
			}
			tx := types.Txs{
				Hash:            &reqParams.TransactionId,
				Address:         reqParams.Address,
				Status:          "0", //未确认
				TransactionType: reqParams.TransactionType,
			}
			err = db.InsertTxs(ctx, txn, tx)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil && strings.Contains(err.Error(), "over limit") {
		return util.JSONResponse{
			Code: http.StatusOK,
			JSON: TransactionInsertResp{
				RetCode: "1",
				Message: "over limit",
			},
		}
	} else if err != nil && !strings.Contains(err.Error(), "over limit") {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.NotFound("db select or insert err"),
		}
	}
	fmt.Println(address + " ::::insertTransaction:::: " + reqParams.TransactionId + " :::: " + reqParams.TransactionType)
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: TransactionInsertResp{
			RetCode: "0",
			Message: "",
		},
	}
}
