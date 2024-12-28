package main

import (
	"os"

	"crypton/crypton"
)

func main() {
	defer os.Exit(0)
	c := &crypton.Crypton{}
	chain := c.Init()
	defer chain.Database.Close()

	cli := CommandLine{chain}
	cli.Run()
}
