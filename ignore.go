package handy

//Ignore ignores crash if program panics
//go:noinline
func Ignore() {
	recover()
}
