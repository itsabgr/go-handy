package handy

import "encoding"

func MustMarshalBinary(value encoding.BinaryMarshaler) []byte {
	b, err := value.MarshalBinary()
	Throw(err)
	return b
}
