package main

import (
	"entysquare/enty-tron-backend/routing"
	"entysquare/enty-tron-backend/storage"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// check duplicated process
	//pid.PassOrPanic()
	//conf := new(conf.Conf).GetConf()
	db, err := storage.NewDatabase()
	if err != nil {
		print("db err:", err)
	}
	if err != nil {
		panic(err)
	}
	routers := mux.NewRouter()
	routing.Setup(routers, db)
	err = http.ListenAndServe("0.0.0.0:9000", routers)
	if err != nil {
		panic("error")
	}
}
