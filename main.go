package main

import (
	"fmt"
	"os"

	"github.com/cycloidio/terracognita/cmd"
)

func main() {
	fmt.Println("aaa")
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
