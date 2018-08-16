package proof

import (
	"fine/blockchain/block"
	"math/big"
	"crypto/sha256"
	"fine/blockchain/util"
	"bytes"
	"math"
	"fmt"
)

type ProofOfWork struct {
	
	Block *block.Block
	
	Target *big.Int //目标值
}

const TargetBits = 24

func NewProofOfWork(block *block.Block) *ProofOfWork{

	Target:=big.NewInt(1)
	Target.Lsh(Target,uint(256-TargetBits))

	pow:=ProofOfWork{Block:block,Target:Target}

	return &pow
}

func (pow *ProofOfWork)Prepare(nonce int64)[]byte  {

	b:=pow.Block

	tmp:=[][]byte{
		util.IntToByte(b.Version),
		b.PrevBlockHash,
		b.MerKelRoot,
		util.IntToByte(b.TimeStamp),
		util.IntToByte(TargetBits),
		util.IntToByte(nonce),
		b.Data,
	}
	data:=bytes.Join(tmp,[]byte{})

	return data
}

func (pow *ProofOfWork)Run()(int64,[]byte){

	var hash [32]byte
	var nonce int64
	var hashInt big.Int
	nonce = 0
	fmt.Printf("target hash:%x\n",pow.Target.Bytes())
	for nonce<math.MaxInt64{
		hash=sha256.Sum256(pow.Prepare(nonce))
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.Target)==-1{
			fmt.Printf("found nonce:%d,%x\n",nonce,hash)

			break
		}else {
			nonce++
		}
	}
	return nonce,hash[:]
	
}

func (pow *ProofOfWork)IsValid()bool{

	var hashInt big.Int
	data:=pow.Prepare(pow.Block.Nonce)
	hash:=sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(pow.Target) == -1
}