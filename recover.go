package handy

//go:noinline
func Recover(fn func()) (recovered interface{}) {
	defer func() {
		recovered = recover()
	}()
	fn()
	return
}
