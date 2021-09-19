package manager

import (
	"fmt"
	"log"
	"time"

	v1 "github.com/ylallemant/panopticon/pkg/api/v1"
	"github.com/ylallemant/panopticon/pkg/chronos"
	"github.com/ylallemant/panopticon/pkg/server/cache"
)

func ObserveUsers(timestamp time.Time, interval time.Duration, users []*v1.User, host *v1.Host, trackedProcesses []*v1.ClassifiedProcess, cache cache.Cache) ([]*v1.UserDailyReport, error) {
	reports := make([]*v1.UserDailyReport, 0)

	for _, user := range users {
		if !cache.IsUserLimited(user.Identifier) {
			// user is not restircted
			continue
		}

		report, err := ObserveUser(timestamp, interval, user, host, trackedProcesses, cache)
		if err != nil {
			return reports, err
		}

		reportLog := fmt.Sprintf(`
daily report for %s:
  creation: %s
  total screen time: %s
  applications:
`,
			user.GetIdentifier(),
			chronos.TimeFromInt64(report.GetTimestamp()),
			chronos.DurationFromInt64(report.GetTotalDuration()),
		)

		for _, application := range report.GetApplications() {
			reportLog += fmt.Sprintf("    %s:\n", application.GetIdentifier())

			for _, duration := range application.GetDurations() {
				reportLog += fmt.Sprintf("      %s: %s (start time %s)\n",
					duration.GetIdentifier(),
					chronos.DurationFromInt64(duration.GetDuration()),
					chronos.TimeFromInt64(duration.GetStartTime()))
			}
		}

		reportLog += fmt.Sprint("  devices:\n")
		for _, device := range report.GetDevices() {
			reportLog += fmt.Sprintf("    %s: %s (start time %s)\n",
				device.GetIdentifier(),
				chronos.DurationFromInt64(device.GetDuration()),
				chronos.TimeFromInt64(device.GetStartTime()))
		}

		log.Print(reportLog)

		reports = append(reports, report)
	}

	return reports, nil
}

func ObserveUser(timestamp time.Time, interval time.Duration, user *v1.User, host *v1.Host, trackedProcesses []*v1.ClassifiedProcess, cache cache.Cache) (*v1.UserDailyReport, error) {
	deviceUserID, found := host.UserIdByIdentifier(user.Identifier)

	if !found {
		return nil, nil
	}

	report, err := FindReport(user, cache)
	if err != nil {
		return nil, err
	}

	// update application usage
	userProcesses := report.FilterProcesses(deviceUserID, trackedProcesses)
	longestDuration := int64(0)

	for _, process := range userProcesses {
		application := report.FindApplication(process.Application.Name)
		device := application.ByDevice(host.Name)

		if device.Duration == 0 {
			device.Duration = process.GetProcess().GetUptime()
			application.Duration = application.Duration + device.Duration
		}

		device.Duration = device.Duration + int64(interval)
		application.Duration = application.Duration + int64(interval)

		if longestDuration < application.Duration {
			longestDuration = application.Duration
		}
	}

	// update device usage
	device := report.FindDevice(host.GetName())

	if device.Duration == 0 {
		device.Duration = longestDuration
	}

	device.Duration = device.Duration + int64(interval)

	return report, nil
}

func FindReport(user *v1.User, cache cache.Cache) (*v1.UserDailyReport, error) {
	report, err := cache.GetUserDailyReport(user.Identifier)
	if err != nil {
		return nil, err
	}

	if report == nil {
		report = NewReport(user.Identifier)
		cache.PutUserDailyReport(report.Identifier, report)
	}

	if !reportActive(report) {
		// TODO: archive old report
		report = NewReport(user.Identifier)
		cache.PutUserDailyReport(report.Identifier, report)
	}

	return report, nil
}

func NewReport(identifier string) *v1.UserDailyReport {
	report := new(v1.UserDailyReport)

	report.Identifier = identifier
	report.Timestamp = chronos.TimestampNano()

	report.Applications = make([]*v1.ApplicationCumulatedTime, 0)
	report.Devices = make([]*v1.CumulatedTime, 0)

	report.TotalDuration = 0

	return report
}

func reportActive(report *v1.UserDailyReport) bool {
	timestamp := chronos.TimeFromInt64(report.Timestamp)
	now := time.Now().UTC()

	return now.Year() == timestamp.Year() &&
		now.Month() == timestamp.Month() &&
		now.Day() == timestamp.Day()
}
