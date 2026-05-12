package format

import (
	"encoding/binary"
	"io"
)

func Write(w io.Writer, encryptedSecret []byte, salt []byte, decryptedLength uint64, chunkSize uint64) error {
	h := header{
		FormatID:              ID,
		EncapsulatedKeyLength: uint64(len(encryptedSecret)),
		SaltSize:              uint64(len(salt)),
		DecryptedLength:       decryptedLength,
		ChunkSize:             chunkSize,
	}

	_, err := w.Write(MagicNumber)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.BigEndian, h)
	if err != nil {
		return err
	}
	_, err = w.Write(encryptedSecret)
	if err != nil {
		return err
	}
	_, err = w.Write(salt)
	if err != nil {
		return err
	}

	return nil
}
