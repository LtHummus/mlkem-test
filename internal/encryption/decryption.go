package encryption

import (
	"crypto/mlkem"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/chacha20poly1305"

	"mlkem-test/internal/format"
)

func decryptChunks(r io.Reader, w io.Writer, key []byte, plaintextBytes uint64, chunkSize uint64) error {
	buf := make([]byte, chunkSize+chacha20poly1305.Overhead)
	var chunkNum uint64
	var bytesRead uint64

	numChunks := (plaintextBytes + chunkSize - 1) / chunkSize

	enc, err := chacha20poly1305.NewX(key)
	if err != nil {
		return err
	}

	for {
		chunkNum++
		n, err := io.ReadFull(r, buf)
		if err != nil && !errors.Is(err, io.EOF) && !errors.Is(err, io.ErrUnexpectedEOF) {
			return err
		}
		lastChunk := numChunks == chunkNum
		data := buf[:n]

		nonce := make([]byte, chacha20poly1305.NonceSizeX)
		binary.BigEndian.PutUint64(nonce, chunkNum)
		aad := binary.BigEndian.AppendUint64(nil, chunkNum)
		if lastChunk {
			aad = append(aad, 0x01)
		} else {
			aad = append(aad, 0x00)
		}

		plaintext, err := enc.Open(nil, nonce, data, aad)
		if err != nil {
			return err
		}
		bytesRead += uint64(n - chacha20poly1305.Overhead)

		_, err = w.Write(plaintext)
		if err != nil {
			return err
		}

		if lastChunk {
			break
		}
	}

	if bytesRead != plaintextBytes {
		return fmt.Errorf("did not read correct number of bytes (expected %d, got %d)", plaintextBytes, bytesRead)
	}

	return nil
}

func Decrypt(dkBytes []byte, r io.Reader, w io.Writer) error {
	dk, err := mlkem.NewDecapsulationKey768(dkBytes)
	if err != nil {
		return err
	}

	encryptedSecret, salt, plainextSize, chunkSize, err := format.Read(r)
	if err != nil {
		return err
	}

	sharedSecret, err := dk.Decapsulate(encryptedSecret)
	if err != nil {
		return err
	}

	key, err := deriveKey(sharedSecret, salt)
	if err != nil {
		return err
	}

	err = decryptChunks(r, w, key, plainextSize, chunkSize)
	if err != nil {
		return err
	}

	return nil
}
