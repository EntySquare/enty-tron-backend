package routing

import (
	"database/sql"
	"entysquare/enty-tron-backend/pkg/jsonerror"
	"entysquare/enty-tron-backend/pkg/util"
	"entysquare/enty-tron-backend/storage"
	"entysquare/enty-tron-backend/storage/sqlutil"
	"entysquare/enty-tron-backend/storage/types"
	"fmt"
	"github.com/tidwall/gjson"
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
	address := strings.ToUpper(gjson.Get(string(reqBody), "address").String())
	err = sqlutil.WithTransaction(db.Db, func(txn *sql.Tx) error {
		addr, err := db.SelectAddressByAddress(ctx, nil, address)
		if err != nil {
			return err
		}
		if addr == nil {
			addrNew := types.Address{
				Address: address,
				Types:   "0",
			}
			err = db.InsertAddress(ctx, txn, addrNew)
			if err != nil {
				return err
			}
		} else {
			if addr.Types != "0" {
				return fmt.Errorf("over limit")
			}
		}
		return nil
	})
	if err != nil && err.Error() != "over limit" {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.NotFound("db select or insert err"),
		}
	} else if err != nil && err.Error() == "over limit" {
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
