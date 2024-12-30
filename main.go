package main

import (
	"fmt"
	"time"
)

func main() {

	bc := NewBlockchain(2)
	bc.AddBlock("First block")
	time.Sleep(2 * time.Second)
	bc.AddBlock("Second block")
	fmt.Println(bc)
}
