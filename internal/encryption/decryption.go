package encryption

import (
	"crypto/mlkem"
	"encoding/binary"
	"errors"
	"io"

	"golang.org/x/crypto/chacha20poly1305"

	"mlkem-test/internal/format"
)

func decryptChunks(r io.Reader, w io.Writer, key []byte) error {
	buf := make([]byte, DefaultChunkSize+chacha20poly1305.Overhead)
	var chunkNum uint64 = 0

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
		lastChunk := errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF)
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

		_, err = w.Write(plaintext)
		if err != nil {
			return err
		}

		if lastChunk {
			break
		}
	}

	return nil
}

func Decrypt(dkBytes []byte, r io.Reader, w io.Writer) error {
	dk, err := mlkem.NewDecapsulationKey768(dkBytes)
	if err != nil {
		return err
	}

	encryptedSecret, salt, err := format.Read(r)
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

	err = decryptChunks(r, w, key)
	if err != nil {
		return err
	}

	return nil
}
