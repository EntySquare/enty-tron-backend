package pid

import (
	"fmt"
	"testing"
)

func TestCheckProRunning(t *testing.T) {
	fmt.Println(CheckProRunning("goland"))
	fmt.Println(GetPid("goland"))
}

func TestProcess(t *testing.T) {
	getCurProcessName()
}
