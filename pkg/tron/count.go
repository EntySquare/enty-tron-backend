package tron

import (
	"fmt"
	"time"
)

func CheckTime() bool {
	timeBegin := BEGIN_TIME
	t1, err := time.Parse(LAYOUT_TIME, timeBegin)
	if err != nil {
		fmt.Println("err:", err)
		panic("time check err")
	}
	timeBegin2 := "2022-05-20 17:00:00"
	t3, err := time.Parse(LAYOUT_TIME, timeBegin2)
	if err != nil {
		fmt.Println("err:", err)
		panic("time check err")
	}
	t2 := time.Now()
	fmt.Println("t1:" + t1.Format(LAYOUT_TIME))
	fmt.Println("t2:" + t2.Format(LAYOUT_TIME))
	fmt.Println("t3:" + t3.Format(LAYOUT_TIME))
	fmt.Println("::::::::::::::::::::::::::::")
	fmt.Println("t2 after t1:")
	fmt.Println(t2.After(t1))
	fmt.Println("::::::::::::::::::::")
	fmt.Println("t2 Before t1:")
	fmt.Println(t2.Before(t1))
	fmt.Println("::::::::::::::::::::")
	fmt.Println("t2 after t3:")
	fmt.Println(t2.After(t3))
	fmt.Println("::::::::::::::::::::")
	fmt.Println("t2 Before t3:")
	fmt.Println(t2.Before(t3))

	fmt.Println("::::::::::::::::::::")
	return t2.After(t1)
}
