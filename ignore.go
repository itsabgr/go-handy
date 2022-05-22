package handy

//go:noinline
func Recover() interface{} {
	return recover()
}
