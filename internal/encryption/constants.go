package encryption

import (
	"crypto/sha256"
	"io"

	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/hkdf"
)

const (
	Info             = "mlkem-test-key | CHACHA20-POLY1305"
	DefaultChunkSize = 64 * 1024 // 64 kb
)

func deriveKey(secret []byte, salt []byte) ([]byte, error) {
	keyStream := hkdf.New(sha256.New, secret, salt, []byte(Info))

	chachaKey := make([]byte, chacha20poly1305.KeySize)
	_, err := io.ReadFull(keyStream, chachaKey)
	if err != nil {
		return nil, err
	}

	return chachaKey, nil
}
