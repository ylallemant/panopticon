package manager

import (
	"time"

	v1 "github.com/ylallemant/panopticon/pkg/api/v1"
	"github.com/ylallemant/panopticon/pkg/server/cache"
)

func Coerce(timestamp time.Time, interval time.Duration, host *v1.Host, trackedUsers []*v1.User, trackedProcesses []*v1.ClassifiedProcess, cache cache.Cache) error {
	_, err := ObserveUsers(timestamp, interval, trackedUsers, host, trackedProcesses, cache)
	if err != nil {
		return err
	}

	return nil
}
