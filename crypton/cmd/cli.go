package main

import (
	"crypton/crypton"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

type (
	CommandLine struct {
		crypton *crypton.Crypton
	}
)

func (cli CommandLine) Cmd() {
	fmt.Println("Usage: ")
	fmt.Println("add -block BLOCK_DATA - add a block to the chain")
	fmt.Println("print - Prints the blocks in the chain")
}

func (cli CommandLine) ValidateArgs() {
	if len(os.Args) < 2 {
		cli.Cmd()
		runtime.Goexit()
	}
}

func (cli CommandLine) AddBlock(data string) {
	cli.crypton.AddBlock(data)
	fmt.Println("Block Added!")
}

func (cli CommandLine) Print() {
	iter := cli.crypton.Iterator()
	for {
		block := iter.Next()

		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := crypton.NewProof(block)
		fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli CommandLine) Run() {
	cli.ValidateArgs()

	add := flag.NewFlagSet("add", flag.ExitOnError)
	print := flag.NewFlagSet("print", flag.ExitOnError)
	data := add.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := add.Parse(os.Args[2:])
		crypton.HandleError(err)
	case "print":
		err := print.Parse(os.Args[2:])
		crypton.HandleError(err)
	default:
		cli.Cmd()
		runtime.Goexit()
	}

	if add.Parsed() {
		if *data == "" {
			add.Usage()
			runtime.Goexit()
		}

		cli.AddBlock(*data)
	}

	if print.Parsed() {
		cli.Print()
	}
}
