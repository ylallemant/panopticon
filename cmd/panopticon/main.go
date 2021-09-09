package main

import (
	"github.com/ylallemant/panopticon/pkg/cli"
	"log"
)

func main() {
	if err := cli.Command().Execute(); err != nil {
		log.Fatalf("error during execution: %v", err)
	}
}
