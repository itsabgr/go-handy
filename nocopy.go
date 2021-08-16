package handy

type NoCopy Empty

func (_ NoCopy) Lock()   {}
func (_ NoCopy) Unlock() {}
