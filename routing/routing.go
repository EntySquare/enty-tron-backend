package routing

import (
	"entysquare/enty-tron-backend/pkg/stream"
	"entysquare/enty-tron-backend/storage"
	"github.com/gorilla/mux"
)

// Setup configures the given mux with sync-server listeners

func Setup(
	csMux *mux.Router, accountDB *storage.Database, cli *stream.Client,
) {
	r0mux := csMux.PathPrefix("/r0").Subrouter()
	println(r0mux)
	////注册Business账户 <none safe>
	//r0mux.Handle("/business/pay/generate",
	//	httputil.MakeExternalAPI("generate", func(req *http.Request) util.JSONResponse {
	//		return generate(req, accountDB)
	//	}),
	//).Methods(http.MethodPost, http.MethodOptions)

	//r0mux.Handle("/user/userInfo",
	//	httputil.MakeAuthAPI("userInfo", accountDB, func(req *http.Request, account *types.Account) util.JSONResponse {
	//		return userInfo(req, accountDB, account)
	//	}),
	//).Methods(http.MethodPost, http.MethodOptions)

}
