package handy

type NoCopy Empty

//go:noinline
func (*NoCopy) Lock() {}

//go:noinline
func (*NoCopy) Unlock() {}
