package blockchain

import (
	"bytes"
	"encoding/binary"
	"log"
	"math/big"
)

// Take the data from the Block
// Create a counter (nonce) which starts at 0
// create a hash of the data plus the counter
// Check the hash to see if meets a set of requirements
// The first few bytes must contain 0s

// blockchain a nivel de mercado , em produção possuem um algoritmo que
// vai incrementando aos poucos a dificuldade.
const difficulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))
	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(difficulty)),
		},
		[]byte{},
	)
	return data
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)

	}
	return buff.Bytes()
}
