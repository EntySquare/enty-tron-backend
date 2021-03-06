package routing

import (
	"entysquare/enty-tron-backend/pkg/tron"
	"entysquare/enty-tron-backend/pkg/util"
	"entysquare/enty-tron-backend/storage"
	"net/http"
	"sync"
)

var resourceLock = func() sync.Mutex {
	return sync.Mutex{}
}()

func insertTransaction(
	req *http.Request, db *storage.Database,
) util.JSONResponse {
	//resourceLock.Lock()
	//defer resourceLock.Unlock()
	//bodyIo := req.Body
	//ctx := req.Context()
	//reqBody, err := ioutil.ReadAll(bodyIo)
	//if err != nil {
	//	return util.JSONResponse{
	//		Code: http.StatusForbidden,
	//		JSON: jsonerror.NotFound("io can not been read successfully"),
	//	}
	//}
	//reqParams := &TransactionInsertReq{}
	//err = json.Unmarshal(reqBody, reqParams)
	//if err != nil {
	//	println(err)
	//	return util.JSONResponse{
	//		Code: http.StatusForbidden,
	//		JSON: jsonerror.Unknown("Transaction unmarshal error"),
	//	}
	//}
	//address := reqParams.Address
	//err = sqlutil.WithTransaction(db.Db, func(txn *sql.Tx) error {
	//	tb, nft, err := db.SelectAllSold(ctx, txn)
	//	if err != nil {
	//		return err
	//	}
	//	if reqParams.TransactionType == tron.TYPE_TB && tb == tron.TBLIMIT {
	//		return fmt.Errorf("over limit")
	//	}
	//	if reqParams.TransactionType == tron.TYPE_TB && nft == tron.NFTLIMIT {
	//		return fmt.Errorf("over limit")
	//	}
	//	addr, err := db.SelectAddressByAddress(ctx, nil, address)
	//	if err != nil {
	//		return err
	//	}
	//	if addr == nil {
	//		return fmt.Errorf("address not exist")
	//	} else {
	//		txs, err := db.SelectTxs(ctx, txn, reqParams.TransactionId)
	//		if txs != nil {
	//			return fmt.Errorf("has been insert")
	//		}
	//		tx := types.Txs{
	//			Hash:            &reqParams.TransactionId,
	//			Address:         reqParams.Address,
	//			Status:          tron.TXS_WAITING_FOR_CHAIN_CONFIRM, //未确认
	//			TransactionType: reqParams.TransactionType,
	//		}
	//		err = db.InsertTxs(ctx, txn, tx)
	//		if err != nil {
	//			return err
	//		}
	//	}
	//	return nil
	//})
	//if err != nil && strings.Contains(err.Error(), "over limit") {
	//	fmt.Println("over limit")
	//	return util.JSONResponse{
	//		Code: http.StatusOK,
	//		JSON: TransactionInsertResp{
	//			RetCode: tron.OVERLIMIT,
	//			Message: "over limit",
	//		},
	//	}
	//} else if err != nil && !strings.Contains(err.Error(), "over limit") {
	//	fmt.Println("limit")
	//	return util.JSONResponse{
	//		Code: http.StatusForbidden,
	//		JSON: jsonerror.NotFound("db select or insert err"),
	//	}
	//}
	//go tron.ScanTron(db, address, reqParams.TransactionType)
	//fmt.Println(address + " ::::insertTransaction:::: " + reqParams.TransactionId + " :::: " + reqParams.TransactionType)
	//return util.JSONResponse{
	//	Code: http.StatusOK,
	//	JSON: TransactionInsertResp{
	//		RetCode: tron.SUCCESS,
	//		Message: "",
	//	},
	//}
	return util.JSONResponse{
		Code: http.StatusOK,
		JSON: TransactionInsertResp{
			RetCode: tron.OVERLIMIT,
			Message: "over limit",
		},
	}
}
