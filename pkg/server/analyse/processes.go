package analyse

import (
	v1 "github.com/ylallemant/panopticon/pkg/api/v1"
)

func filterParentProcessesByMappedUsers(host *v1.Host, report *v1.HostProcessReportRequest) ([]*v1.Process, error) {
	list := make([]*v1.Process, 0)

	for _, process := range report.Report.Processes {
		if process.PPID > 1 {
			continue
		}

		for _, user := range host.UserMapping {
			if process.GetUserID() != user.Id {
				continue
			}

			list = append(list, process)
			/*log.Printf(
				"PID: %d\tPPID: %d\tChilds: %d\t=> %d\t%s => %s",
				process.GetPID(),
				process.GetPPID(),
				process.GetChilds(),
				process.GetUserID(),
				process.GetUser(),
				process.GetCommand(),
			)*/
		}
	}

	return list, nil
}
