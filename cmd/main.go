package main

import (
	"fmt"
	"os"

	"github.com/Leocanela279/cli-password-manager/internal/vault"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: vault <command>")
		return
	}

	command := os.Args[1]

	switch command {
	case "init":
		err := vault.Init()

		if err != nil {
			fmt.Println(err)
			return
		}
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("usage: vault add <service> <username> <password>")
			return
		}
	default:
		fmt.Println("unknown command")
	}
}
