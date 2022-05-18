package main

import (
	"entysquare/enty-tron-backend/conf"
	"entysquare/enty-tron-backend/storage"
)

func main() {
	// check duplicated process
	//pid.PassOrPanic()
	conf := new(conf.Conf).GetConf()
	db, err := storage.NewDatabase()
	if err != nil {
		print("db err:", err)
	}
	if err != nil {
		panic(err)
	}
	println(conf.Postgres.Dbname)
	println(db.Db.Ping())
}
