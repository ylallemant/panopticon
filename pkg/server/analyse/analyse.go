package analyse

import (
	"log"

	v1 "github.com/ylallemant/panopticon/pkg/api/v1"
	"github.com/ylallemant/panopticon/pkg/chronos"
	"github.com/ylallemant/panopticon/pkg/server/cache"
	"github.com/ylallemant/panopticon/pkg/server/classifier"
	"github.com/ylallemant/panopticon/pkg/server/manager"
)

func Analyse(report *v1.HostProcessReportRequest, cache cache.Cache, classifier classifier.Classifier) error {
	log.Printf("------\nAnalyse host %s with %d reported processes",
		report.Report.GetHostname(), len(report.GetReport().GetProcesses()))

	host, err := Host(report, cache)
	if err != nil {
		return err
	}
	identifyNewTrackedUsers(host, cache)

	trackedUsers, err := getTrackedUsers(host, cache)
	if err != nil {
		return err
	}

	log.Printf("teacked users: %d", len(trackedUsers))

	userParentProcesses, err := filterParentProcessesByMappedUsers(host, report)
	if err != nil {
		return err
	}

	log.Printf("teacked user parent process count: %d", len(userParentProcesses))

	trackedProcesses, err := classifyProcesses(host, userParentProcesses, classifier)
	if err != nil {
		return err
	}

	log.Printf("teacked processes: %d", len(trackedProcesses))

	timestamp := chronos.TimeFromInt64(report.Report.Timestamp)
	interval := chronos.DurationFromInt64(report.Report.Interval)

	return manager.Coerce(timestamp, interval, host, trackedUsers, trackedProcesses, cache)
}
