package util

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func CalculateInt64(x int64, y int64, operator string) (i int64) {
	switch operator {
	case "add":
		a := big.NewInt(x)
		b := big.NewInt(y)
		z := a.Add(a, b)
		i := z.Int64()
		return i
	case "sub":
		a := big.NewInt(x)
		b := big.NewInt(y)
		z := a.Sub(a, b)
		i := z.Int64()
		return i
	case "mul":
		a := big.NewInt(x)
		b := big.NewInt(y)
		z := a.Mul(a, b)
		i := z.Int64()
		return i
	case "div":
		a := big.NewInt(x)
		b := big.NewInt(y)
		z := a.Div(a, b)
		i := z.Int64()
		return i
	}
	return i
}
func CalculateString(x string, y string, operator string) (i string, definedErr *MessageError) {
	if x == "" {
		x = "0"
	}
	if y == "" {
		y = "0"
	}
	a := new(big.Float)
	b := new(big.Float)
	xbf, xok := a.SetString(x)
	if !xok {
		definedErr := NewMsgError(4, "error in trans string into big float")
		return "", definedErr
	}
	ybf, yok := b.SetString(y)
	if !yok {
		definedErr := NewMsgError(4, "error in trans string into big float")
		return "", definedErr
	}
	switch operator {
	case "add":
		ai, xErr := strconv.Atoi(x)
		bi, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			definedErr = NewMsgError(4, "error in trans string into int")
		}
		i := strconv.Itoa(ai + bi)
		return i, definedErr
	case "sub":
		ai, xErr := strconv.Atoi(x)
		bi, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			definedErr = NewMsgError(4, "error in trans string into int")
		}
		if ai < bi {
			return "", NewMsgError(0, "wrong order")
		}
		i := strconv.Itoa(ai - bi)
		return i, definedErr
	case "mul":
		ai, xErr := strconv.Atoi(x)
		bi, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			definedErr = NewMsgError(4, "error in trans string into int")
		}
		i := strconv.Itoa(ai * bi)
		return i, definedErr
	case "div":
		ai, xErr := strconv.Atoi(x)
		bi, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			definedErr = NewMsgError(4, "error in trans string into int")
		}
		if bi == 0 {
			return "", NewMsgError(0, "the denominator is zero")
		}
		i := strconv.Itoa(ai / bi)
		return i, definedErr
	case "cmp":
		ai, xErr := strconv.Atoi(x)
		bi, yErr := strconv.Atoi(y)
		if xErr != nil || yErr != nil {
			definedErr = NewMsgError(4, "error in trans string into int")
		}
		if ai > bi {
			i = "1"
		} else if ai < bi {
			i = "-1"
		} else {
			i = "0"
		}
		return i, definedErr
	case "addBigFU":
		ibf := xbf.Add(xbf, ybf)
		i = ibf.Text('f', 18)
		return i, definedErr
	case "subBigFU":
		if xbf.Cmp(ybf) < 0 {
			definedErr := NewMsgError(0, "wrong order")
			return "", definedErr
		}
		ibf := xbf.Sub(xbf, ybf)
		i = ibf.Text('f', 18)
		return i, definedErr
	case "addBigFH":
		ibf := xbf.Add(xbf, ybf)
		i = ibf.Text('f', 2)
		return i, definedErr
	case "cmpBigF":
		i = strconv.Itoa(xbf.Cmp(ybf))
		return i, definedErr
	case "divBigF":
		ibf := xbf.Quo(xbf, ybf)
		i = ibf.Text('f', 18) //保留6位小数
		return i, definedErr
	case "mulBigF":
		ibf := xbf.Mul(xbf, ybf)
		i = ibf.Text('f', 18) //保留6位小数
		return i, definedErr
	case "subBigFH":
		if xbf.Cmp(ybf) < 0 {
			definedErr := NewMsgError(0, "wrong order")
			return "", definedErr
		}
		ibf := xbf.Sub(xbf, ybf)
		i = ibf.Text('f', 2)
		return i, definedErr
	}

	return i, definedErr
}
func Digit(x string, operator string) (i string, definedErr *MessageError) {
	unit := new(big.Float)
	a := new(big.Float)
	bf, ok := a.SetString(x)
	if !ok {
		definedErr := NewMsgError(4, "error in trans string into bigint")
		return "", definedErr
	}
	switch operator {
	case "div18":
		unit.SetString("1000000000000000000")
		bf.Quo(bf, unit)
		i = bf.Text('f', 18)
		return i, definedErr
	case "div18-f7":
		unit.SetString("1000000000000000000")
		bf.Quo(bf, unit)
		i = bf.Text('f', 7)
		return i, definedErr
	case "div18-f4":
		unit.SetString("1000000000000000000")
		bf.Quo(bf, unit)
		i = bf.Text('f', 4)
		return i, definedErr
	case "div2":
		unit.SetString("100")
		bf.Quo(bf, unit)
		i = bf.Text('f', 2)
		return i, definedErr
	case "mul2":
		unit.SetString("100")
		bf.Mul(bf, unit)
		i = bf.Text('f', 0)
		return i, definedErr
	case "mul18":
		unit.SetString("1000000000000000000")
		bf.Mul(bf, unit)
		i = bf.Text('f', 0)
		return i, definedErr
	}

	return i, definedErr
}

/*
	换算 *18位
*/
func StrToFil18(num string) (string, error) {

	numCompany := 0

	//检查点符号 数量
	arr := strings.Split(num, ".")
	if len(arr) >= 3 {
		return "", errors.New("num err 001")
	}

	//检查是否有负数
	var check = strings.Index(num, "-")
	if check != -1 {
		return "", errors.New("num err 002")
	}

	//扣掉符号 .
	numStr := strings.Replace(num, ".", "", -1)

	var index = strings.Index(num, ".")
	if index >= 1 {
		numCompany = len(numStr) - index
		fmt.Println("小数单位", numCompany)
	}
	decimal := 18 - numCompany
	for decimal > 0 {
		decimal--
		numStr = numStr + "0"
	}

	fmt.Println("字符串结果：", numStr)
	fmt.Println(numStr)
	return numStr, nil
}

func StrNanoFILToFilStr(nanoFilStr string, num string) string {
	float, err := strconv.ParseFloat(nanoFilStr, 64)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%."+num+"f", float/1000000000.0)
}

// CalculateBigIntString calculate two string as big int in using string format
func CalculateBigIntString(a, b, operator string) string {
	as, ok := new(big.Int).SetString(a, 10)
	if !ok {
		panic("big int string calculate error")
	}
	bs, ok := new(big.Int).SetString(b, 10)
	if !ok {
		panic("big int string calculate error")
	}
	res := new(big.Int)
	switch operator {
	case "add":
		res.Add(as, bs)
	case "minus":
		res.Sub(as, bs)
	case "cmp":
		res.Set(as).Cmp(bs)
	}
	return res.String()
}
func CmpValue(walletValue string, orderValue string, percent float64) (bool, error) {
	denoFloat, err := strconv.ParseFloat(walletValue, 64)
	if err != nil {
		return false, err
	}
	orderFloat, err := strconv.ParseFloat(orderValue, 64)
	if err != nil {
		return false, err
	}
	if denoFloat*percent >= orderFloat {
		return true, nil
	}
	return false, nil
}
