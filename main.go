package main

import (
	"fmt"

	"github.com/sebastian-j-ibanez/ledger"
)

func main() {
	l := ledger.NewLedger()

	println()
	for i := range 3 {
		msg := fmt.Sprintf("This is the %d entry in the ledger.", i)
		fmt.Printf("Appending node %d\n", i+1)
		err := l.AddNode([]byte(msg))
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Println("\nHashes:")
	for _, node := range l.GetNodes() {
		fmt.Printf("%d\n", node.Hash)
	}

	valid, err := l.ValidateLedger()
	if err != nil {
		panic(err.Error())
	}

	if valid {
		println("\nLedger contents are cryptographically valid.\n")
	} else {
		println("\nLedger contents are NOT cryptographically valid.\n")
	}
}
