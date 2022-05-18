package util

import (
	"fmt"
	"math/big"
	"strconv"
)

/*
	币换算
	data 换算数据 例：0x0000000000000000000000000000000000000000000000000000000000989680
	decimal 保留小数位
	proportion 换算比例
*/
func CoinCount(data string, decimal string, proportion float64) float64 {
	data1, _ := new(big.Float).SetString(data)
	i, _ := new(big.Float).Quo(data1, big.NewFloat(proportion)).Float64()
	i = Decimal(i, decimal)
	return i
}
func Decimal(value float64, i string) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%."+i+"f", value), 64)
	return value
}

//字符串 是否包含在素组内
func StrToArrForTrue(str string, arr []string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}
