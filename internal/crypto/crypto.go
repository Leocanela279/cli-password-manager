package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func MakeHash(masterkey string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(masterkey), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return hash, nil
}

func CompareHash(masterkey string, pass string) error {

	return bcrypt.CompareHashAndPassword([]byte(masterkey), []byte(pass))
}
