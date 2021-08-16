package handy

func Throw(errors ...interface{}) {
	for _, error := range errors {
		if error != nil {
			panic(error)
		}
	}
}
