package main

import (
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	fingerprint := newFingerprintFromPkgutil(filename)
	fmt.Println(fingerprint)
}
