package main

import (
	"fmt"
	"os"

	"github.com/GLAD-DEV/dn/cmd"
	"github.com/GLAD-DEV/dn/config"
)

func main() {
	err := config.CreateDir()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	_ = cmd.Execute()
}
