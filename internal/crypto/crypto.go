package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
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

func deriveKey(secret string) []byte {
	// Pad or truncate to 32 bytes for AES-256
	key := make([]byte, 32)
	copy(key, []byte(secret))
	return key
}

func Encrypt(secret, pass string) (string, error) {
	key := deriveKey(secret)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(pass), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func Decrypt(secret, encryptedPass string) (string, error) {
	key := deriveKey(secret)
	
	cipherText, err := base64.StdEncoding.DecodeString(encryptedPass)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", errors.New("decryption failed: wrong key or corrupted data")
	}

	return string(plainText), nil
}
