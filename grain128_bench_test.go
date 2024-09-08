package grain128

import (
	"testing"
)

// Benchmark for keystream generation
func BenchmarkGrain128Keystream(b *testing.B) {
	key := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F}
	iv := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0A, 0x0B}
	cipher, err := NewGrain128(key)
	if err != nil {
		b.Fatalf("Failed to initialize Grain128 cipher: %v", err)
	}
	cipher.IVSetup(iv)

	keystream := make([]byte, 16)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cipher.KeystreamBytes(keystream)
	}
}

// Benchmark for encryption
func BenchmarkGrain128Encrypt(b *testing.B) {
	key := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F}
	iv := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0A, 0x0B}

	plaintext := []byte("SECRET MESSAGE!!")
	ciphertext := make([]byte, len(plaintext))

	cipher, err := NewGrain128(key)
	if err != nil {
		b.Fatalf("Failed to initialize Grain128 cipher: %v", err)
	}
	cipher.IVSetup(iv)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cipher.XORKeyStream(ciphertext, plaintext)
	}
}

// Benchmark for decryption
func BenchmarkGrain128Decrypt(b *testing.B) {
	key := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F}
	iv := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0A, 0x0B}

	plaintext := []byte("SECRET MESSAGE!!")
	ciphertext := make([]byte, len(plaintext))
	decrypted := make([]byte, len(plaintext))

	cipher, err := NewGrain128(key)
	if err != nil {
		b.Fatalf("Failed to initialize Grain128 cipher: %v", err)
	}
	cipher.IVSetup(iv)

	// First encrypt
	cipher.XORKeyStream(ciphertext, plaintext)

	// Benchmark decryption
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Note: IV must be reset before decryption
		cipher.IVSetup(iv)
		cipher.XORKeyStream(decrypted, ciphertext)
	}
}
