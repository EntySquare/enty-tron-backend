package routing

import (
	"entysquare/enty-tron-backend/pkg/httputil"
	"entysquare/enty-tron-backend/pkg/util"
	"entysquare/enty-tron-backend/storage"
	"github.com/gorilla/mux"
	"net/http"
)

// Setup configures the given mux with sync-server listeners

func Setup(
	csMux *mux.Router, accountDB *storage.Database,
) {
	r0mux := csMux.PathPrefix("/r0").Subrouter()
	//检查地址使用 <none safe>
	r0mux.Handle("/tron/address/checkAddress",
		httputil.MakeExternalAPI("checkAddress", func(req *http.Request) util.JSONResponse {
			return checkAddress(req, accountDB)
		}),
	).Methods(http.MethodPost, http.MethodOptions)

	r0mux.Handle("/tron/transaction/insertTransaction",
		httputil.MakeExternalAPI("insertTransaction", func(req *http.Request) util.JSONResponse {
			return insertTransaction(req, accountDB)
		}),
	).Methods(http.MethodPost, http.MethodOptions)

	r0mux.Handle("/tron/address/queryCoinLimit",
		httputil.MakeExternalAPI("queryCoinLimit", func(req *http.Request) util.JSONResponse {
			return queryCoinLimit(req, accountDB)
		}),
	).Methods(http.MethodPost, http.MethodOptions)

	r0mux.Handle("/tron/address/confirmLimit",
		httputil.MakeExternalAPI("confirmLimit", func(req *http.Request) util.JSONResponse {
			return confirmLimit(req, accountDB)
		}),
	).Methods(http.MethodPost, http.MethodOptions)
	//r0mux.Handle("/user/userInfo",
	//	httputil.MakeAuthAPI("userInfo", accountDB, func(req *http.Request, account *types.Account) util.JSONResponse {
	//		return userInfo(req, accountDB, account)
	//	}),
	//).Methods(http.MethodPost, http.MethodOptions)

}
