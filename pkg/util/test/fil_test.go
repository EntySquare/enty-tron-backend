package main

import (
	"encoding/hex"
	"fmt"
	//"github.com/ethereum/go-ethereum/crypto"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-crypto"
	"testing"
)

func TestAddress(t *testing.T) {
	sk, err := crypto.GenerateKey()
	fmt.Printf("privateKey: %v\n", sk)
	fmt.Printf("privateKey: %s  len:%v err: %v\n", hex.EncodeToString(sk), len(hex.EncodeToString(sk)), err)

	publicKey := crypto.PublicKey(sk)
	fmt.Printf("publicKey : %v\n", publicKey)
	fmt.Printf("publicKey : %s  len:%v err: %v\n", hex.EncodeToString(publicKey), len(hex.EncodeToString(publicKey)), err)

	addr, err := address.NewSecp256k1Address(publicKey)
	fmt.Printf("address   : %v\n", addr)
	fmt.Printf("address   : err:%v\n", err)

}
