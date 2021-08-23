package handy

import "time"

//ZeroTime returns a zero time
func ZeroTime() *time.Time {
	return &time.Time{}
}

//IsZeroTime checks aTime is zero
func IsZeroTime(aTime *time.Time) bool {
	return time.Time{} == *aTime
}

//Time is a nullable time
type Time struct {
	Time time.Time
	//IsFill is true when is not nil
	IsFill bool
}
