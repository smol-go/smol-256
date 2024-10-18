package main

import (
	"fmt"
	"os"

	"github.com/themillenniumfalcon/smol-256/algo"
)

func main() {
	sha := algo.Sha256{}

	var str string

	if len(os.Args) < 2 {
		str = ""
	} else {
		str = os.Args[1]
	}

	hash := sha.Hash(str)

	fmt.Println(algo.ToHexString(hash))
}
