package algo

import (
	"bytes"
	"fmt"
	"math"
)

type Sha256 struct {
	Digest [32]byte
}

func (sha256 *Sha256) Hash(str string) {
	msg := []byte(str)
	msg = bytes.Join([][]byte{msg, {0x80}}, []byte{})

	l := float64(len(msg)) / float64(64)
	N := uint32(math.Ceil(l))

	fmt.Println("msg is", N)
}
