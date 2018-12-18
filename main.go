package main

import (
	"os"

	"github.com/nfvsap/natsbeat/cmd"

	_ "github.com/nfvsap/natsbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
