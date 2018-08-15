package blockchain

import b "fine/blockchain/block"

type BlockChain struct {

	Blocks []*b.Block


}

func NewBlockChain() *BlockChain{

	block:=b.NewGenesisBlock()
	return &BlockChain{Blocks:[]*b.Block{block}}
}

func (bc *BlockChain) AddBlock(data string) {
	block:=b.NewBlock(data,bc.Blocks[len(bc.Blocks)-1].Hash)
	bc.Blocks = append(bc.Blocks, block)

}

