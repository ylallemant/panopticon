package v1

import (
	"github.com/ylallemant/panopticon/pkg/chronos"
)

func (r *UserDailyReport) FindDevice(hostname string) *CumulatedTime {
	for _, report := range r.Devices {
		if report.Identifier == hostname {
			return report
		}
	}

	device := new(CumulatedTime)
	device.Identifier = hostname
	device.StartTime = chronos.TimestampNano()
	device.Duration = 0

	r.Devices = append(r.Devices, device)

	return device
}

func (r *UserDailyReport) FindApplication(applicationName string) *ApplicationCumulatedTime {
	for _, application := range r.Applications {
		if application.Identifier == applicationName {
			return application
		}
	}

	application := new(ApplicationCumulatedTime)
	application.Identifier = applicationName
	application.Duration = 0
	application.Durations = make([]*CumulatedTime, 0)

	r.Applications = append(r.Applications, application)

	return application
}

func (r *UserDailyReport) FilterProcesses(hostUserId int32, processes []*ClassifiedProcess) []*ClassifiedProcess {
	filtered := make([]*ClassifiedProcess, 0)

	for _, process := range processes {
		if process.Process.GetUserID() == hostUserId {
			filtered = append(filtered, process)
		}
	}

	return filtered
}

func (r *ApplicationCumulatedTime) ByDevice(hostname string) *CumulatedTime {
	for _, duration := range r.Durations {
		if duration.Identifier == hostname {
			return duration
		}
	}

	device := new(CumulatedTime)
	device.Identifier = hostname
	device.StartTime = chronos.TimestampNano()
	device.Duration = 0

	r.Durations = append(r.Durations, device)

	return device
}
