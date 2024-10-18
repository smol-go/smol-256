package algo

import (
	"fmt"
)

type Sha256 struct {
	Digest [32]byte
}

func (sha256 *Sha256) Hash(str string) {
	msg := []byte(str)
	fmt.Println("msg is", msg)
}
