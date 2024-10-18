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

	M := make([][]uint32, N)
	for i := range M {
		M[i] = make([]uint32, 16)
	}

	for i := range M {
		for j := 0; j < 16; j++ {
			index := i*64 + j*4

			for k := 0; k < 4; k++ {
				if index+k < len(msg) {
					M[i][j] |= (uint32(msg[index+k]) << uint32(8*(3-k)))
				}
			}
		}
	}

	fmt.Println("msg is", M)
}
