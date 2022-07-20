package handy

import "bytes"

func Equal(a, b any) bool {
	switch (a).(type) {
	case []byte:
		return bytes.Equal(a.([]byte), b.([]byte))
	}
	return a == b
}
