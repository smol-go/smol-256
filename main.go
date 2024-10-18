package main

import (
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

	sha.Hash(str)
}
