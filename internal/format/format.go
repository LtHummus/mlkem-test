package format

const ID uint8 = 1

var MagicNumber = []byte("MLKEMTEST")

type header struct {
	FormatID              uint8
	EncapsulatedKeyLength uint64
	SaltSize              uint64
}
