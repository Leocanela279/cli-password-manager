package crypto

import (
	"testing"
)

func TestMakeHash(t *testing.T) {
	hash, err := MakeHash("testpassword")
	if err != nil {
		t.Fatalf("MakeHash failed: %v", err)
	}

	if len(hash) == 0 {
		t.Error("Expected non-empty hash")
	}

	// Hash should be different each time due to salt
	hash2, _ := MakeHash("testpassword")
	if string(hash) == string(hash2) {
		t.Error("Same password should produce different hashes (due to salt)")
	}
}

func TestCompareHash(t *testing.T) {
	password := "testpassword123"
	
	hash, err := MakeHash(password)
	if err != nil {
		t.Fatalf("MakeHash failed: %v", err)
	}

	// Correct password should match
	err = CompareHash(string(hash), password)
	if err != nil {
		t.Error("Correct password should match hash")
	}

	// Wrong password should not match
	err = CompareHash(string(hash), "wrongpassword")
	if err == nil {
		t.Error("Wrong password should not match hash")
	}
}

func TestCompareHashEmptyPassword(t *testing.T) {
	hash, _ := MakeHash("password")
	
	err := CompareHash(string(hash), "")
	if err == nil {
		t.Error("Empty password should not match hash")
	}
}

func TestEncrypt(t *testing.T) {
	secret := "this-is-a-32-byte-secret-key-here!" // exactly 32 bytes
	password := "mypassword"

	encrypted, err := Encrypt(secret, password)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	if encrypted == password {
		t.Error("Encrypted password should be different from original")
	}

	if len(encrypted) == 0 {
		t.Error("Encrypted password should not be empty")
	}

	// Encryption should produce different output each time (due to random IV)
	encrypted2, _ := Encrypt(secret, password)
	if encrypted == encrypted2 {
		t.Error("Same password should produce different ciphertext (due to random IV)")
	}
}

func TestEncryptDifferentSecrets(t *testing.T) {
	password := "mypassword"
	secret1 := "this-is-a-32-byte-secret-key-one!"
	secret2 := "this-is-a-32-byte-secret-key-two!"

	encrypted1, _ := Encrypt(secret1, password)
	encrypted2, _ := Encrypt(secret2, password)

	if encrypted1 == encrypted2 {
		t.Error("Same password with different keys should produce different ciphertext")
	}
}

func TestEncryptEmptyPassword(t *testing.T) {
	secret := "this-is-a-32-byte-secret-key-here!"
	
	encrypted, err := Encrypt(secret, "")
	if err != nil {
		t.Fatalf("Encrypt failed for empty password: %v", err)
	}

	if len(encrypted) == 0 {
		t.Error("Encrypted empty string should not be empty")
	}
}

func TestEncryptLongPassword(t *testing.T) {
	secret := "this-is-a-32-byte-secret-key-here!"
	longPassword := "this-is-a-very-long-password-that-exceeds-16-bytes"
	
	encrypted, err := Encrypt(secret, longPassword)
	if err != nil {
		t.Fatalf("Encrypt failed for long password: %v", err)
	}

	if len(encrypted) == 0 {
		t.Error("Encrypted long password should not be empty")
	}
}

func TestEncryptSpecialCharacters(t *testing.T) {
	secret := "this-is-a-32-byte-secret-key-here!"
	password := "p@ssw0rd!#$%^&*()_+-=[]{}|;':\",./<>?"

	encrypted, err := Encrypt(secret, password)
	if err != nil {
		t.Fatalf("Encrypt failed for special characters: %v", err)
	}

	if len(encrypted) == 0 {
		t.Error("Encrypted password with special characters should not be empty")
	}
}

func TestEncryptUnicode(t *testing.T) {
	secret := "this-is-a-32-byte-secret-key-here!"
	password := "密码测试🔐"

	encrypted, err := Encrypt(secret, password)
	if err != nil {
		t.Fatalf("Encrypt failed for unicode: %v", err)
	}

	if len(encrypted) == 0 {
		t.Error("Encrypted unicode password should not be empty")
	}
}

func TestDecrypt(t *testing.T) {
	secret := "this-is-a-32-byte-secret-key-here!"
	password := "mypassword123"

	encrypted, err := Encrypt(secret, password)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	decrypted, err := Decrypt(secret, encrypted)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}

	if decrypted != password {
		t.Errorf("Expected decrypted password '%s', got '%s'", password, decrypted)
	}
}

func TestDecryptDifferentPasswords(t *testing.T) {
	secret := "this-is-a-32-byte-secret-key-here!"
	
	encrypted1, _ := Encrypt(secret, "password1")
	encrypted2, _ := Encrypt(secret, "password2")

	decrypted1, _ := Decrypt(secret, encrypted1)
	decrypted2, _ := Decrypt(secret, encrypted2)

	if decrypted1 != "password1" {
		t.Errorf("Expected 'password1', got '%s'", decrypted1)
	}

	if decrypted2 != "password2" {
		t.Errorf("Expected 'password2', got '%s'", decrypted2)
	}
}

func TestDecryptWrongSecret(t *testing.T) {
	secret1 := "this-is-a-32-byte-secret-key-one!"
	secret2 := "this-is-a-32-byte-secret-key-two!"
	password := "mysecretpassword"

	encrypted, _ := Encrypt(secret1, password)
	
	_, err := Decrypt(secret2, encrypted)
	if err == nil {
		t.Error("Decrypt with wrong secret should fail")
	}
}

func TestDecryptEmptyPassword(t *testing.T) {
	secret := "this-is-a-32-byte-secret-key-here!"
	
	encrypted, _ := Encrypt(secret, "")
	decrypted, err := Decrypt(secret, encrypted)
	
	if err != nil {
		t.Fatalf("Decrypt failed for empty password: %v", err)
	}

	if decrypted != "" {
		t.Errorf("Expected empty string, got '%s'", decrypted)
	}
}

func BenchmarkMakeHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeHash("benchmark-password")
	}
}

func BenchmarkEncrypt(b *testing.B) {
	secret := "this-is-a-32-byte-secret-key-here!"
	for i := 0; i < b.N; i++ {
		Encrypt(secret, "benchmark-password")
	}
}

func BenchmarkCompareHash(b *testing.B) {
	hash, _ := MakeHash("benchmark-password")
	hashStr := string(hash)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CompareHash(hashStr, "benchmark-password")
	}
}