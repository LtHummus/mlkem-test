package encryption

import (
	"crypto/hkdf"
	"crypto/sha256"

	"golang.org/x/crypto/chacha20poly1305"
)

const (
	Info             = "mlkem-test-key | CHACHA20-POLY1305"
	DefaultChunkSize = 64 * 1024 // 64 kb
)

func deriveKey(secret []byte, salt []byte) ([]byte, error) {
	return hkdf.Key(sha256.New, secret, salt, Info, chacha20poly1305.KeySize)
}
