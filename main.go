package main

import (
	"fmt"
	"github.com/gasperio/gasper/cmd"
	"os"
)

func main() {
	if err := cmd.Root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
