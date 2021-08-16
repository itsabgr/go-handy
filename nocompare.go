package handy

//NoCompare prevents struct to be compared
type NoCompare struct {
	NoCompare [0]func()
}
