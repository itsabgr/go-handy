package handy

import "fmt"

//Cast [T any] casts interface{} to T
func Cast[T any](t interface{}) T {
	switch t.(type) {
	case T:
		return t.(T)
	}
	panic(fmt.Errorf("unsupported type %T", t))
}
