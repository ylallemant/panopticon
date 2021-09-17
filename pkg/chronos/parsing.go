package chronos

import (
	"time"
)

func TimestampNano() int64 {
	return int64(time.Now().UTC().UnixNano())
}

func TimeFromInt64(timestamp int64) time.Time {
	return time.Unix(0, timestamp).UTC()
}

func DurationFromInt64(timestamp int64) time.Duration {
	return time.Duration(timestamp)
}
