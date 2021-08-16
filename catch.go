package handy

func Catch(fn func(recovered interface{})) {
	recovered := recover()
	if recovered != nil {
		fn(recovered)
	}
}
