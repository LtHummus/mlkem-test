package encryption

import (
	"bufio"
	"crypto/mlkem"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"

	"golang.org/x/crypto/chacha20poly1305"

	"mlkem-test/internal/format"
)

func encryptChunks(r io.Reader, w io.Writer, chunkSize int, plaintextLength uint64, key []byte) error {
	buf := make([]byte, chunkSize)
	var chunkNum uint64
	var bytesEncrypted uint64

	enc, err := chacha20poly1305.NewX(key)
	if err != nil {
		return err
	}

	r = bufio.NewReader(r)

	for {
		chunkNum++
		n, err := io.ReadFull(r, buf)
		if err != nil && !errors.Is(err, io.EOF) && !errors.Is(err, io.ErrUnexpectedEOF) {
			return err
		}
		bytesEncrypted += uint64(n)
		lastChunk := bytesEncrypted >= plaintextLength
		data := buf[:n]

		nonce := make([]byte, chacha20poly1305.NonceSizeX)
		binary.BigEndian.PutUint64(nonce, chunkNum)
		aad := binary.BigEndian.AppendUint64(nil, chunkNum)
		if lastChunk {
			aad = append(aad, 0x01)
		} else {
			aad = append(aad, 0x00)
		}

		ciphertext := enc.Seal(nil, nonce, data, aad)

		_, err = w.Write(ciphertext)
		if err != nil {
			return err
		}

		if lastChunk {
			break
		}
	}

	return nil
}

func Encrypt(ekBytes []byte, plaintextLength uint64, fileToEncrypt io.Reader, encryptFile io.Writer) error {
	ek, err := mlkem.NewEncapsulationKey768(ekBytes)
	if err != nil {
		return err
	}

	secret, encryptedSecret := ek.Encapsulate()

	salt := make([]byte, 32)
	_, err = io.ReadFull(rand.Reader, salt)
	if err != nil {
		return err
	}

	chachaKey, err := deriveKey(secret, salt)
	if err != nil {
		return err
	}

	err = format.Write(encryptFile, encryptedSecret, salt, plaintextLength, DefaultChunkSize)
	if err != nil {
		return err
	}

	err = encryptChunks(fileToEncrypt, encryptFile, DefaultChunkSize, plaintextLength, chachaKey)
	if err != nil {
		return err
	}

	return nil
}
