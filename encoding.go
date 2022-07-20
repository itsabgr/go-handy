package handy

import "encoding"

func MustMarshalBinary(value encoding.BinaryMarshaler) []byte {
	b, err := value.MarshalBinary()
	Throw(err)
	return b
}

func MustMarshalText(value encoding.TextMarshaler) string {
	txt, err := value.MarshalText()
	Throw(err)
	return txt
}
