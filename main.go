package main

import (
	"github.com/bbokka123/nomadcoin/cli"
	"github.com/bbokka123/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
