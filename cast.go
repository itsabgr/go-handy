package handy

import "fmt"

//Cast casts any type to interface{}
func Cast[T any](t T) interface{} {
	switch t.(type) {
	case T:
		return t.(T)
	}
	panic(fmt.Errorf("unsupported type %T", t))
}
