package main

import (
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

var infuraURL = "https://mainnet.infura.io/v3/a5c80040e89b490c9f19bffdc28110e3"

func main() {
	conn, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Something went wee wrong!", err)
	}
}
