package format

import (
	"encoding/binary"
	"fmt"
	"io"
	"slices"
)

func Read(r io.Reader) ([]byte, []byte, error) {
	magicByte := make([]byte, len(MagicNumber))
	_, err := io.ReadFull(r, magicByte)
	if err != nil {
		return nil, nil, err
	}

	if !slices.Equal(magicByte, MagicNumber) {
		return nil, nil, fmt.Errorf("invalid magic number")
	}

	var h header
	err = binary.Read(r, binary.BigEndian, &h)
	if err != nil {
		return nil, nil, err
	}

	if h.EncapsulatedKeyLength > 16*1024 {
		return nil, nil, fmt.Errorf("invalid encapsulated key length")
	}

	encryptedSecret := make([]byte, h.EncapsulatedKeyLength)
	_, err = io.ReadFull(r, encryptedSecret)
	if err != nil {
		return nil, nil, err
	}

	if h.SaltSize > 16*1024 {
		return nil, nil, fmt.Errorf("invalid salt length")
	}

	salt := make([]byte, h.SaltSize)
	_, err = io.ReadFull(r, salt)
	if err != nil {
		return nil, nil, err
	}

	return encryptedSecret, salt, nil
}
