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
