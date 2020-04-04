package main

import (
	"fmt"
	"github.com/bxsmart/go-hdwallet"
	"log"
)

func main() {
	seed, _ := go_hdwallet.NewSeed()
	wallet, err := go_hdwallet.NewFromSeed(seed)
	if err != nil {
		log.Fatal(err)
	}

	path := go_hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947

	path = go_hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account, err = wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x8230645aC28A4EdD1b0B53E7Cd8019744E9dD559
}
