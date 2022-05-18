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
)

func insertTransaction(
	req *http.Request, db *storage.Database,
) util.JSONResponse {
	bodyIo := req.Body
	ctx := req.Context()
	reqBody, err := ioutil.ReadAll(bodyIo)
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.NotFound("io can not been read successfully"),
		}
	}
	reqParams := TransactionInsertReq{}
	err = json.Unmarshal(reqBody, reqParams)
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.Unknown("Transaction unmarshal error"),
		}
	}
	address := reqParams.Address
	err = sqlutil.WithTransaction(db.Db, func(txn *sql.Tx) error {
		addr, err := db.SelectAddressByAddress(ctx, nil, address)
		if err != nil {
			return err
		}
		if addr == nil {
			return fmt.Errorf("address not exist")
		} else {
			tx := types.Txs{
				Hash:    &reqParams.TransactionId,
				Address: reqParams.Address,
				Status:  "0", //未确认
			}
			err = db.InsertTxs(ctx, txn, tx)
			if err != nil {
				return err
			}
			addr.Types = reqParams.TransactionType
			err = db.UpdateAddressById(ctx, txn, *addr)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: TransactionInsertResp{
			RetCode: "0",
			Message: "",
		},
	}
}
