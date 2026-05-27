package vault

import (
	"errors"
	"fmt"
	"os"
)

func Init() error {
	err := os.WriteFile("../vault-test/.vault", []byte("{}"), 0644)
	if err != nil {
		return errors.New("error while trying vault init")
	}

	return nil
}

func Add(service string, username string, password string) {
	data, err := os.ReadFile("../../vault-test/.vault")

	if err != nil {
		fmt.Println("error while trying to add into the vault")
		return
	}
	fmt.Printf("data: %v\n", data)
}
