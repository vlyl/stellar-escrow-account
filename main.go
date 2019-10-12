package main

import (
	"github.com/tevino/log"
	"github.com/vlyl/stellar-escrow-account/config"
	"os"
)

func main() {
	log.Info("stellar escrow account project")

	curdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	config.LoadConfig(curdir)
}
