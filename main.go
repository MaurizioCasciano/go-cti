package main

import (
	"fmt"
	"github.com/vulsio/go-cti/commands"
	"os"
)

// @title Go-CTI
// @description This is a CTI application
// @version 1.0
// @host localhost:1329
// @BasePath /
func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}
