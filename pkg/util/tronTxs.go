package util

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

func CheckTransaction(hash string) (bool, error) {
	resp, err := http.Get("https://apilist.tronscan.org/api/transaction-info?hash=" + hash)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var contractRet = gjson.Get(string(b), "contractRet").String()
	fmt.Println(contractRet)
	if contractRet == "SUCCESS" {
		return true, nil
	}
	if contractRet == "OUT_OF_ENERGY" || contractRet == "" {
		return false, nil
	}
	return false, errors.New("undetermined")
}
