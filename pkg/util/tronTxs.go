package util

import (
	"errors"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
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
	if contractRet == "SUCCESS" {
		var amount_str = gjson.Get(string(b), "tokenTransferInfo.amount_str").String()
		var to_address = gjson.Get(string(b), "tokenTransferInfo.to_address").String()
		//钱包不对
		if strings.ToUpper(to_address) != strings.ToUpper("TNNU5FtQT8nfdyogZuYy1pRo7ToBe7bvEH") &&
			strings.ToUpper(to_address) != strings.ToUpper("THPT8T7ikAJZzsuZ4rD54qVbXEUnGgAAXg") {
			return false, nil
		}
		//金额不对
		if amount_str != "100000000" {
			return false, nil
		}
		return true, nil
	}
	if contractRet == "OUT_OF_ENERGY" || contractRet == "" {
		return false, nil
	}
	return false, errors.New("undetermined")
}
