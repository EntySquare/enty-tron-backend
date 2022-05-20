package routing

import (
	"database/sql"
	"encoding/json"
	"entysquare/enty-tron-backend/conf"
	"entysquare/enty-tron-backend/pkg/jsonerror"
	"entysquare/enty-tron-backend/pkg/tron"
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
		RetCode:  tron.SUCCESS,
		TbLimit:  tron.UNCOMMITTED,
		NftLimit: tron.UNCOMMITTED,
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
				resp.TbLimit = tron.WAITING_FOR_PAYMENT
				if txs != nil {
					if txs.Status == tron.WAITING_FOR_PAYMENT {
						resp.TbLimit = tron.WAITING_FOR_DELIVERY
					} else if txs.Status == tron.TXS_ADMIN_CONFIRM {
						resp.TbLimit = tron.DELIVERED
					}
				}

			}
			if addr.Nft == "1" {
				txs, err := db.SelectTxsByAddressAndType(ctx, txn, address, "2")
				if err != nil {
					return err
				}
				resp.NftLimit = tron.WAITING_FOR_PAYMENT
				if txs != nil {
					if txs.Status == tron.TXS_CHAIN_CONFIRM {
						resp.NftLimit = tron.WAITING_FOR_DELIVERY
					} else if txs.Status == tron.TXS_ADMIN_CONFIRM {
						resp.NftLimit = tron.DELIVERED
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("checkAddress fail")
		return util.JSONResponse{
			Code: http.StatusForbidden,
			JSON: jsonerror.NotFound("db select or insert err"),
		}
	}
	//fmt.Println("checkAddress")
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: resp,
	}
}
func queryCoinLimit(
	req *http.Request, db *storage.Database,
) util.JSONResponse {
	////bodyIo := req.Body
	//ctx := req.Context()
	////reqBody, err := ioutil.ReadAll(bodyIo)
	////if err != nil {
	////	return util.JSONResponse{
	////		Code: http.StatusForbidden,
	////		JSON: jsonerror.NotFound("io can not been read successfully"),
	////	}
	////}
	//_, nft, err := db.SelectAllSold(ctx, nil)
	//if err != nil {
	//	fmt.Println("queryCoinLimit fail")
	//	return util.JSONResponse{
	//		Code: http.StatusForbidden,
	//		JSON: jsonerror.NotFound("db select  err"),
	//	}
	//}
	////fmt.Println("queryCoinLimit")
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: QueryCoinLimitResp{
			RetCode:        tron.SUCCESS,
			TbHasBeenSold:  200,
			NftHasBeenSold: 200,
		},
	}
}
func confirmLimit(req *http.Request, db *storage.Database,
) util.JSONResponse {
	////resourceLock.Lock()
	////defer resourceLock.Unlock()
	////bodyIo := req.Body
	////ctx := req.Context()
	////reqBody, err := ioutil.ReadAll(bodyIo)
	////if err != nil {
	////	return util.JSONResponse{
	////		Code: http.StatusForbidden,
	////		JSON: jsonerror.NotFound("io can not been read successfully"),
	////	}
	////}
	////reqParams := &ConfirmLimitReq{}
	////err = json.Unmarshal(reqBody, reqParams)
	////if err != nil {
	////	println(err)
	////	return util.JSONResponse{
	////		Code: http.StatusForbidden,
	////		JSON: jsonerror.Unknown("Transaction unmarshal error"),
	////	}
	////}
	////address := reqParams.Address
	////if address == "" {
	////	return util.JSONResponse{
	////		Code: http.StatusForbidden,
	////		JSON: jsonerror.NotFound("address is null"),
	////	}
	////}
	////err = sqlutil.WithTransaction(db.Db, func(txn *sql.Tx) error {
	////	tb, nft, err := db.SelectAllSold(ctx, txn)
	////	if err != nil {
	////		return err
	////	}
	////	if reqParams.TransactionType == tron.TYPE_TB && tb == tron.TBLIMIT {
	////		return fmt.Errorf("over limit")
	////	}
	////	if reqParams.TransactionType == tron.TYPE_NFT && nft == tron.NFTLIMIT {
	////		return fmt.Errorf("over limit")
	////	}
	////	addr, err := db.SelectAddressByAddress(ctx, nil, address)
	////	if err != nil {
	////		return err
	////	}
	////	if addr == nil {
	////		return fmt.Errorf("address not exist")
	////	} else {
	////		if reqParams.TransactionType == tron.TYPE_TB {
	////			if addr.Tb == "1" {
	////				return fmt.Errorf("has been sold")
	////			}
	////			if addr.Nft == "1" {
	////				txs, err := db.SelectTxsByAddressAndType(ctx, txn, address, "2")
	////				if err != nil {
	////					return err
	////				}
	////				if txs == nil {
	////					return fmt.Errorf("nft locked")
	////				}
	////			}
	////			addr.Tb = "1"
	////
	////		} else if reqParams.TransactionType == tron.TYPE_NFT {
	////			if addr.Nft == "1" {
	////				return fmt.Errorf("has been sold")
	////			}
	////			if addr.Tb == "1" {
	////				txs, err := db.SelectTxsByAddressAndType(ctx, txn, address, "1")
	////				if err != nil {
	////					return err
	////				}
	////				if txs == nil {
	////					return fmt.Errorf("tb locked")
	////				}
	////			}
	////			addr.Nft = "1"
	////		}
	////		err = db.UpdateAddressById(ctx, txn, *addr)
	////		if err != nil {
	////			return err
	////		}
	////	}
	////	return nil
	////})
	////if err != nil && strings.Contains(err.Error(), "has been sold") {
	////	return util.JSONResponse{
	////		Code: http.StatusOK,
	////		JSON: ConfirmLimitResp{
	////			RetCode: tron.SOLD,
	////			Message: "have been sold",
	////		},
	////	}
	////} else if err != nil && strings.Contains(err.Error(), "locked") {
	////	return util.JSONResponse{
	////		Code: http.StatusOK,
	////		JSON: ConfirmLimitResp{
	////			RetCode: tron.LOCKED,
	////			Message: "locked",
	////		},
	////	}
	////} else if err != nil && strings.Contains(err.Error(), "over limit") {
	////	return util.JSONResponse{
	////		Code: http.StatusForbidden,
	////		JSON: ConfirmLimitResp{
	////			RetCode: tron.IDOOVERLIMIT,
	////			Message: "over limit",
	////		},
	////	}
	////} else if err != nil && !strings.Contains(err.Error(), "has been sold") && !strings.Contains(err.Error(), "locked") && !strings.Contains(err.Error(), "over limit") {
	////	return util.JSONResponse{
	////		Code: http.StatusForbidden,
	////		JSON: jsonerror.Unknown(" confirm limit error"),
	////	}
	////}
	////go tron.UnlockLimit(db, address)
	//fmt.Println(address + " ::::confirmLimit:::: " + reqParams.TransactionType)
	//return util.JSONResponse{
	//	Code: http.StatusOK,
	//	JSON: ConfirmLimitResp{
	//		RetCode: tron.SUCCESS,
	//		Message: "",
	//	},
	//}
	return util.JSONResponse{
		Code: http.StatusForbidden,
		JSON: ConfirmLimitResp{
			RetCode: tron.IDOOVERLIMIT,
			Message: "over limit",
		},
	}
}

func returnLimit(req *http.Request, db *storage.Database,
) util.JSONResponse {
	//bodyIo := req.Body
	//ctx := req.Context()
	//reqBody, err := ioutil.ReadAll(bodyIo)
	//if err != nil {
	//	return util.JSONResponse{
	//		Code: http.StatusForbidden,
	//		JSON: jsonerror.NotFound("io can not been read successfully"),
	//	}
	//}
	//reqParams := &ReturnLimitReq{}
	//err = json.Unmarshal(reqBody, reqParams)
	//if err != nil {
	//	println(err)
	//	return util.JSONResponse{
	//		Code: http.StatusForbidden,
	//		JSON: jsonerror.Unknown("Transaction unmarshal error"),
	//	}
	//}
	//address := reqParams.Address
	//if address == "" {
	//	return util.JSONResponse{
	//		Code: http.StatusForbidden,
	//		JSON: jsonerror.NotFound("address is null"),
	//	}
	//}
	//err = sqlutil.WithTransaction(db.Db, func(txn *sql.Tx) error {
	//	addr, err := db.SelectAddressByAddress(ctx, txn, address)
	//	if err != nil {
	//		return err
	//	}
	//	var count = 0
	//	if addr.Nft == "1" {
	//		txs, err := db.SelectTxsByAddressAndType(ctx, txn, addr.Address, "2")
	//		if err != nil {
	//			return err
	//		}
	//		if txs == nil {
	//			count += 1
	//			fmt.Println(addr.Address + " ::::return limit:::: 2")
	//		}
	//	}
	//	if addr.Tb == "1" {
	//		txs, err := db.SelectTxsByAddressAndType(ctx, txn, addr.Address, "1")
	//		if err != nil {
	//			return err
	//		}
	//		if txs == nil {
	//			count += 2
	//
	//			fmt.Println(addr.Address + " ::::return limit:::: 1")
	//		}
	//	}
	//	if count == 3 {
	//		addr.Tb = "0"
	//		addr.Nft = "0"
	//	} else if count == 2 {
	//		addr.Tb = "0"
	//	} else if count == 1 {
	//		addr.Nft = "0"
	//	}
	//	err = db.UpdateAddressById(ctx, txn, *addr)
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//})
	//if err != nil && !strings.Contains(err.Error(), "has been sold") {
	//	return util.JSONResponse{
	//		Code: http.StatusForbidden,
	//		JSON: jsonerror.Unknown(" confirm limit error"),
	//	}
	//} else if err != nil && strings.Contains(err.Error(), "has been sold") {
	//	return util.JSONResponse{
	//		Code: http.StatusAccepted,
	//		JSON: ReturnLimitResp{
	//			RetCode: tron.SOLD,
	//			Message: "has been sold",
	//		},
	//	}
	//}
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: ReturnLimitResp{
			RetCode: tron.OVERLIMIT,
			Message: "over limit",
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
			if tbtxs != nil && tbtxs.Status == tron.TXS_CHAIN_CONFIRM {
				addrsStatus := AddressStatus{
					Address:         addr.Address,
					TransactionId:   *tbtxs.Hash,
					TransactionType: tbtxs.TransactionType,
					Status:          tron.TXS_CHAIN_CONFIRM,
				}
				list = append(list, addrsStatus)

			} else if tbtxs != nil && tbtxs.Status == tron.TXS_CHAIN_CONFIRM {
				addrsStatus := AddressStatus{
					Address:         addr.Address,
					TransactionId:   *tbtxs.Hash,
					TransactionType: tbtxs.TransactionType,
					Status:          tron.TXS_CHAIN_CONFIRM,
				}
				list = append(list, addrsStatus)
			}

			nfttxs, err := db.SelectTxsByAddressAndType(ctx, txn, addr.Address, "2")
			if err != nil {
				return err
			}
			if nfttxs != nil && nfttxs.Status == tron.TXS_CHAIN_CONFIRM {
				addrsStatus := AddressStatus{
					Address:         addr.Address,
					TransactionId:   *nfttxs.Hash,
					TransactionType: nfttxs.TransactionType,
					Status:          tron.TXS_CHAIN_CONFIRM,
				}
				list = append(list, addrsStatus)

			} else if nfttxs != nil && nfttxs.Status == tron.TXS_ADMIN_CONFIRM {
				addrsStatus := AddressStatus{
					Address:         addr.Address,
					TransactionId:   *nfttxs.Hash,
					TransactionType: nfttxs.TransactionType,
					Status:          tron.TXS_ADMIN_CONFIRM,
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
			RetCode: tron.SUCCESS,
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
			txs.Status = tron.TXS_ADMIN_CONFIRM
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
	fmt.Println("setAddressStatus")
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: SetAddressStatusResp{
			RetCode: tron.SUCCESS,
			Message: "",
		},
	}
}
