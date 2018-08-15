package main

import (
	"fmt"
	"fine/blockchain/blockchain"
)


func main() {

	bc:=blockchain.NewBlockChain()
	bc.AddBlock("A send B 1 BTC")
	bc.AddBlock("B send C 1 BTC")

	for _,block:=range bc.Blocks {

		fmt.Printf("version:%d\n",block.Version)
		fmt.Printf("PrevBlockHash:%x\n",block.PrevBlockHash)
		fmt.Printf("Hash:%x\n",block.Hash)
		fmt.Printf("TimeStamp:%d\n",block.TimeStamp)
		fmt.Printf("Bits:%d\n",block.Bits)
		fmt.Printf("Nonce:%d\n",block.Nonce)
		fmt.Printf("Data:%s\n",block.Nonce)

	}

}
