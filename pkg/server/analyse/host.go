package analyse

import (
	"log"

	v1 "github.com/ylallemant/panopticon/pkg/api/v1"
	"github.com/ylallemant/panopticon/pkg/server/cache"
)

func Host(report *v1.HostProcessReportRequest, cache cache.Cache) (*v1.Host, error) {

	cached, err := cache.GetHost(report.Report.GetHostname())
	if err != nil {
		return nil, err
	}

	if cached == nil {
		cached = new(v1.Host)

		cached.Name = report.Report.Hostname
		cached.Arch = report.Report.Arch
		cached.Os = report.Report.Os
		cached.Admins = []*v1.HostUser{}
		cached.Users = []*v1.HostUser{}
		cached.UserMapping = map[string]*v1.HostUser{}
		cached.Processes = map[string]*v1.ProcessReport{}
		cached.DefaultDurationLimits = []*v1.DailyDurationLimit{}
		cached.DefaultTimeLimits = []*v1.DailyTimeLimit{}
		cached.HasChanged = true

		cache.PutHost(cached.Name, cached)
	}

	cached.HasChanged = syncUsers(cached, report.Report.Processes)

	if cached.HasChanged {
		cache.PutHost(cached.Name, cached)
	}

	return cached, nil
}

func syncUsers(host *v1.Host, processes []*v1.Process) bool {
	changed := false

	for _, process := range processes {
		if !host.UserIdExists(process.GetUserID()) {
			user := new(v1.HostUser)
			user.Id = process.GetUserID()
			user.Name = process.GetUser()

			log.Printf("add user %s (%d) for host %s", user.Name, user.Id, host.Name)

			host.Users = append(host.Users, user)
			changed = true
		}
	}

	return changed
}
