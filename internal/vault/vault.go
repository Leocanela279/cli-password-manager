package vault

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Leocanela279/cli-password-manager/internal/crypto"
	"golang.org/x/term"
)

type Vault struct {
	Username  string           `json:"username"`
	MasterKey string           `json:"masterkey"`
	Entries   map[string]Entry `json:"entries"`
}
type Entry struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Init() error {
	err := os.WriteFile("../vault-test/.vault", []byte("{}"), 0644)
	if err != nil {
		return fmt.Errorf("error while trying vault init: %w", err)
	}

	return nil
}

func Add(service string, username string, password string) error {
	data, err := os.ReadFile("../vault-test/.vault")

	if err != nil {
		return fmt.Errorf("error while trying to add into the vault: %w", err)
	}
	var vault Vault

	err = json.Unmarshal(data, &vault)
	if err != nil {
		return err
	}

	if vault.Entries == nil {
		vault.Entries = make(map[string]Entry)
	}

	vault.Entries[service] = Entry{
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
		return fmt.Errorf("an error ocurred while trying to read your vault: %w", err)
	}

	var vault Vault

	err = json.Unmarshal(data, &vault)

	fmt.Printf("service %s pass: %s\n", service, vault.Entries[service].Password)
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

func List() error {
	data, err := os.ReadFile("../vault-test/.vault")
	if err != nil {
		return fmt.Errorf("failed to read vault file: %w", err)
	}

	var vault Vault

	err = json.Unmarshal(data, &vault)
	if err != nil {
		return fmt.Errorf("failed to parse vault data: %w", err)
	}

	if len(vault.Entries) == 0 {
		fmt.Println("No services saved")
		return nil
	}

	fmt.Println("Available services:")

	for service := range vault.Entries {
		fmt.Println("-", service)
	}

	return nil
}

func Remove(service string) error {
	data, err := os.ReadFile("../vault-test/.vault")
	if err != nil {
		return fmt.Errorf("failed to read vault file: %w", err)
	}

	var vault map[string]Entry

	err = json.Unmarshal(data, &vault)
	if err != nil {
		return fmt.Errorf("failed to parse vault data: %w", err)
	}

	_, exists := vault[service]

	if !exists {
		return fmt.Errorf("service '%s' does not exist", service)
	}

	delete(vault, service)

	updatedData, err := json.MarshalIndent(vault, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode vault data: %w", err)
	}

	err = os.WriteFile("../vault-test/.vault", updatedData, 0644)
	if err != nil {
		return fmt.Errorf("failed to update vault file: %w", err)
	}

	fmt.Println("Service removed successfully")

	return nil
}

func Login() error {
	var username string
	fmt.Print("username:")
	fmt.Scan(&username)

	fmt.Print("password:")
	pass, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	hash, err := crypto.MakeHash(string(pass))

	if err != nil {
		return err
	}

	var vault Vault
	data, err := os.ReadFile("../vault-test/.vault")

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &vault)
	if err != nil {
		return err
	}

	vault.Username = username
	vault.MasterKey = string(hash)

	updatedData, err := json.MarshalIndent(vault, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("../vault-test/.vault", updatedData, 0644)
}
