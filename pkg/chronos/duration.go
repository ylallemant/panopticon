package chronos

import "time"

func ToSeconds(duration time.Duration) int64 {
	return int64(duration / time.Second)
}
