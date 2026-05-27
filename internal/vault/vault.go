package vault

import (
	"encoding/json"
	"errors"
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
