package routing

import (
	"database/sql"
	"encoding/json"
	"entysquare/enty-tron-backend/conf"
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
	if address == "" {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.NotFound("address is null"),
		}
	}
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
				txs, err := db.SelectTxsByAddressAndType(ctx, txn, address, "1")
				if err != nil {
					return err
				}
				if txs != nil {
					if txs.Status == "1" {
						resp.TbLimit = "3"
					} else {
						resp.TbLimit = "2"
					}
				}
				resp.TbLimit = "1"
			}
			if addr.Nft == "1" {
				txs, err := db.SelectTxsByAddressAndType(ctx, txn, address, "2")
				if err != nil {
					return err
				}
				if txs != nil {
					if txs.Status == "1" {
						resp.NftLimit = "3"
					} else {
						resp.NftLimit = "2"
					}
				}
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
	if address == "" {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.NotFound("address is null"),
		}
	}
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
	if err != nil && !strings.Contains(err.Error(), "has been sold") {
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
func listAddressJoined(req *http.Request, db *storage.Database,
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
	reqParams := &ListAddressJoinedReq{}
	err = json.Unmarshal(reqBody, reqParams)
	if err != nil {
		println(err)
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.Unknown("Transaction unmarshal error"),
		}
	}
	var conf = new(conf.Conf).GetConf()
	if reqParams.Username != conf.Admin.Username || reqParams.Password != conf.Admin.Password {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.Unknown("Wrong token "),
		}
	}
	list := make([]AddressStatus, 0)
	err = sqlutil.WithTransaction(db.Db, func(txn *sql.Tx) error {
		addrs, err := db.ListAddressByStatus(ctx, txn)
		if err != nil {
			return err
		}
		for _, addr := range addrs {
			tbtxs, err := db.SelectTxsByAddressAndType(ctx, txn, addr.Address, "1")
			if err != nil {
				return err
			}
			if tbtxs != nil && tbtxs.Status == "1" {
				addrsStatus := AddressStatus{
					Address:         addr.Address,
					TransactionId:   *tbtxs.Hash,
					TransactionType: tbtxs.TransactionType,
					Status:          "1",
				}
				list = append(list, addrsStatus)

			} else if tbtxs != nil && tbtxs.Status == "2" {
				addrsStatus := AddressStatus{
					Address:         addr.Address,
					TransactionId:   *tbtxs.Hash,
					TransactionType: tbtxs.TransactionType,
					Status:          "2",
				}
				list = append(list, addrsStatus)
			}

			nfttxs, err := db.SelectTxsByAddressAndType(ctx, txn, addr.Address, "2")
			if err != nil {
				return err
			}
			if nfttxs != nil && nfttxs.Status == "1" {
				addrsStatus := AddressStatus{
					Address:         addr.Address,
					TransactionId:   *nfttxs.Hash,
					TransactionType: nfttxs.TransactionType,
					Status:          "1",
				}
				list = append(list, addrsStatus)

			} else if nfttxs != nil && nfttxs.Status == "2" {
				addrsStatus := AddressStatus{
					Address:         addr.Address,
					TransactionId:   *nfttxs.Hash,
					TransactionType: nfttxs.TransactionType,
					Status:          "2",
				}
				list = append(list, addrsStatus)
			}
		}
		return nil
	})
	if err != nil && !strings.Contains(err.Error(), "has been sold") {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.Unknown(" list Address error"),
		}
	}
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: ListAddressJoinedResp{
			RetCode: "0",
			List:    list,
		},
	}
}
func setAddressJoined(req *http.Request, db *storage.Database,
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
	reqParams := &SetAddressStatusReq{}
	err = json.Unmarshal(reqBody, reqParams)
	if err != nil {
		println(err)
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.Unknown("SetAddressStatusReq unmarshal error"),
		}
	}
	var conf = new(conf.Conf).GetConf()
	if reqParams.Username != conf.Admin.Username || reqParams.Password != conf.Admin.Password {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.Unknown("Wrong token "),
		}
	}
	err = sqlutil.WithTransaction(db.Db, func(txn *sql.Tx) error {
		txs, err := db.SelectTxsByAddressAndType(ctx, txn, reqParams.AddrStatus.Address, reqParams.AddrStatus.TransactionType)
		if err != nil {
			return err
		}
		if txs != nil {
			txs.Status = "2"
			err = db.UpdateTxsByHash(ctx, txn, *txs)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil && !strings.Contains(err.Error(), "has been sold") {
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.Unknown(" set status  error"),
		}
	}
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: SetAddressStatusResp{
			RetCode: "0",
			Message: "",
		},
	}
}
