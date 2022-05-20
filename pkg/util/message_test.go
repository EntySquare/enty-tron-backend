package util

import (
	"fmt"
	"testing"
)

func TestSendCheckCodeMessage(t *testing.T) {
	_ = SendCheckCodeMessage("13620202901", "006893")
}

func TestEmail(t *testing.T) {
	SingleMail("terilscaub@gmail.com", "184684")
}

func TestEmail2(t *testing.T) {
	fmt.Println(CheckTransaction("c8a138a8caee2ec2f319f1c59c373d6041bf543b1190cf4df849597f20be092c"))
}
