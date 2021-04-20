package structs

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/thaison247/go-blockchain/configs"
	"math"
	"math/big"
	"strconv"
)

var (
	TARGET_BITS = configs.GetConfig().TARGET_BITS
	maxNonce = math.MaxInt64
)

type ProofOfWork struct {
	Block *Block
	Target *big.Int // 0000001000000000000000000000000
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256 - TARGET_BITS)) // target << (256 - target)

	pow := &ProofOfWork{b, target}

	return pow
}

// prepare data to be hash: merge block's data & target_bits & nonce to []bytes
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			[]byte(strconv.FormatInt(int64(pow.Block.Timestamp), 16)),
			[]byte(strconv.FormatInt(int64(TARGET_BITS), 16)),
			[]byte(strconv.FormatInt(int64(nonce), 16)),
		},
		[]byte{},
	)

	return data
}

// hash array of bytes (include block's data & target_bits & nonce)
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.Block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Print("\n\n")

	return nonce, hash[:]
}

