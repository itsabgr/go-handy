package handy

import "io"

//Close closes the closable without error
func Close(closable io.Closer) {
	_ = closable.Close()
}
