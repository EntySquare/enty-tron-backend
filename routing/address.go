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
)

func checkAddress(
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
	reqParams := &CheckAddressReq{}
	err = json.Unmarshal(reqBody, reqParams)
	if err != nil {
		println(err)
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.Unknown("Transaction unmarshal error"),
		}
	}
	address := reqParams.Address
	resp := CheckAddressResp{
		RetCode:  "0",
		TbLimit:  "0",
		NftLimit: "0",
	}
	err = sqlutil.WithTransaction(db.Db, func(txn *sql.Tx) error {
		addr, err := db.SelectAddressByAddress(ctx, nil, address)
		if err != nil {
			return err
		}
		if addr == nil {
			addrNew := types.Address{
				Address: address,
				Nft:     "0",
				Tb:      "0",
			}
			err = db.InsertAddress(ctx, txn, addrNew)
			if err != nil {
				return err
			}
		} else {
			if addr.Tb == "1" {
				resp.TbLimit = "1"
			} else if addr.Nft == "1" {
				resp.NftLimit = "1"
			}
		}
		return nil
	})
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.NotFound("db select or insert err"),
		}
	}

	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: resp,
	}
}
func queryCoinLimit(
	req *http.Request, db *storage.Database,
) util.JSONResponse {
	//bodyIo := req.Body
	ctx := req.Context()
	//reqBody, err := ioutil.ReadAll(bodyIo)
	//if err != nil {
	//	return util.JSONResponse{
	//		Code: http.StatusForbidden,
	//		JSON: jsonerror.NotFound("io can not been read successfully"),
	//	}
	//}
	tb, nft, err := db.SelectAllSold(ctx, nil)
	if err != nil {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.NotFound("db select  err"),
		}
	}
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: QueryCoinLimitResp{
			RetCode:        "0",
			TbHasBeenSold:  tb,
			NftHasBeenSold: nft,
		},
	}
}
func confirmLimit(req *http.Request, db *storage.Database,
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
	reqParams := &ConfirmLimitReq{}
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
			if reqParams.TransactionType == "1" {
				if addr.Tb == "1" {
					return fmt.Errorf("has been sold")
				}
				addr.Tb = "1"
			} else if reqParams.TransactionType == "2" {
				if addr.Nft == "1" {
					return fmt.Errorf("has been sold")
				}
				addr.Nft = "1"
			}
			err = db.UpdateAddressById(ctx, txn, *addr)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil && strings.Contains(err.Error(), "has been sold") {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.Unknown("has been sold"),
		}
	} else if err != nil && !strings.Contains(err.Error(), "has been sold") {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.Unknown(" confirm limit error"),
		}
	}
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: ConfirmLimitResp{
			RetCode: "0",
			Message: "",
		},
	}
}
