package main

import (
	"fmt"
	"os"

  "github.com/fs0414/nodeid-transfer/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
