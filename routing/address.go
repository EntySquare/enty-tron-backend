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
			if reqParams.TransactionType == "1" && addr.Tb == "1" {
				return fmt.Errorf("tb coin over limit")
			} else if reqParams.TransactionType == "2" && addr.Nft == "1" {
				return fmt.Errorf("nft over limit")
			}
		}
		return nil
	})
	if err != nil && !strings.Contains(err.Error(), "over limit") {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.NotFound("db select or insert err"),
		}
	} else if err != nil && strings.Contains(err.Error(), "over limit") {
		return util.JSONResponse{
			Code: http.StatusOK,
			JSON: CheckAddressResp{
				RetCode: "1",
				Message: "over limit",
			},
		}
	}

	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: CheckAddressResp{
			RetCode: "0",
			Message: "",
		},
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
