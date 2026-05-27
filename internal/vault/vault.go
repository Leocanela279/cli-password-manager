package vault

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Entry struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Init() error {
	err := os.WriteFile("../vault-test/.vault", []byte("{}"), 0644)
	if err != nil {
		return errors.New("error while trying vault init")
	}

	return nil
}

func Add(service string, username string, password string) error {
	data, err := os.ReadFile("../vault-test/.vault")

	if err != nil {
		return errors.New("error while trying to add into the vault")
	}
	var vault map[string]Entry

	err = json.Unmarshal(data, &vault)
	if err != nil {
		return err
	}

	vault[service] = Entry{
		Username: username,
		Password: password,
	}

	updatedData, err := json.MarshalIndent(vault, "", "  ")
	if err != nil {
		return err
	}
	os.WriteFile("../vault-test/.vault", updatedData, 0644)
	return nil
}

func Get(service string) error {
	data, err := os.ReadFile("../vault-test/.vault")

	if err != nil {
		return errors.New("an error ocurred while trying to read your vault")
	}

	var vault map[string]Entry

	err = json.Unmarshal(data, &vault)

	fmt.Printf("service %s pass: %s\n", service, vault[service].Password)
	return nil
}

func Help() {
	commands := map[string]string{
		"init": "Initialize your personal vault",
		"add":  "Add a new service with credentials",
		"get":  "Get credentials for a service",
		"list": "List all saved services",
		"help": "Show available commands",
	}

	fmt.Println("Usage:")
	fmt.Println("  vault <command>")
	fmt.Println()

	fmt.Println("Available Commands:")

	for command, description := range commands {
		fmt.Printf("  %-6s: %s\n", command, description)
	}
}
