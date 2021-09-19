package analyse

import (
	"fmt"
	"log"

	v1 "github.com/ylallemant/panopticon/pkg/api/v1"
	"github.com/ylallemant/panopticon/pkg/reporting"
)

func filterParentProcessesByMappedUsers(host *v1.Host, report *v1.HostProcessReportRequest) ([]*v1.Process, error) {
	list := make([]*v1.Process, 0)

	rootProcess := reporting.GetRootProcess(report.Report.GetOs(), report.Report.Processes)
	if rootProcess == nil {
		return list, fmt.Errorf("root process could not be found on %s device %s", report.Report.GetOs(), report.Report.GetHostname())
	}

	for _, process := range report.Report.Processes {
		if process.GetPPID() > rootProcess.GetPID() {
			continue
		}

		reporting.PrintProcess(process)

		for _, user := range host.UserMapping {
			if process.GetUserID() != user.Id {
				continue
			}

			list = append(list, process)
		}
	}

	log.Print("root process:")
	reporting.PrintProcess(rootProcess)

	return list, nil
}
