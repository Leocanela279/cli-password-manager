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
	case "login":
		err := vault.Login()

		if err != nil {
			fmt.Println(err)
			return
		}
	case "add":

		if len(os.Args) < 5 {
			fmt.Println("usage: vault add <service> <username> <password>")
			return
		}

		service := os.Args[2]
		username := os.Args[3]
		password := os.Args[4]
		err := vault.Add(service, username, password)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("vault added!!")

	case "get":
		if len(os.Args) < 3 {
			fmt.Println("usage: vault get <service>")
			return
		}

		service := os.Args[2]

		err := vault.Get(service)

		if err != nil {
			fmt.Println(err)
			return
		}

	case "help":
		vault.Help()
	case "list":
		err := vault.List()

		if err != nil {
			fmt.Println(err)
			return
		}
	case "ls":
		err := vault.List()

		if err != nil {
			fmt.Println(err)
			return
		}
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("usage: vault remove <service>")
			return
		}

		service := os.Args[2]
		err := vault.Remove(service)

		if err != nil {
			fmt.Println(err)
			return
		}

	case "rm":
		if len(os.Args) < 3 {
			fmt.Println("usage: vault rm <service>")
			return
		}

		service := os.Args[2]
		err := vault.Remove(service)

		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		fmt.Println("unknown command")
	}
}
