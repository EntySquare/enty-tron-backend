package util

import "testing"

func TestSendCheckCodeMessage(t *testing.T) {
	_ = SendCheckCodeMessage("13620202901", "006893")
}

func TestEmail(t *testing.T) {
	SingleMail("terilscaub@gmail.com", "184684")
}
