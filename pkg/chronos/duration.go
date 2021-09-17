package chronos

import "time"

func DurationToNanoseconds(duration time.Duration) int64 {
	return duration.Nanoseconds()
}
