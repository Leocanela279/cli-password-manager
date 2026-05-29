package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

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

func Encrypt(secret, pass string) (string, error) {
	block, err := aes.NewCipher([]byte(secret))

	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len([]byte(pass)))
	iv := cipherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", nil
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[:aes.BlockSize], []byte(pass))

	return string(cipherText), nil
}
