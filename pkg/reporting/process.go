package reporting

import (
	"log"
	"regexp"

	v1 "github.com/ylallemant/panopticon/pkg/api/v1"
	"github.com/ylallemant/panopticon/pkg/chronos"
)

var (
	rootProcessRegexps     = map[string]*regexp.Regexp{}
	rootProcessExpressions = map[string]string{
		"linux":  "lxpanel",
		"darwin": "\\/sbin\\/launchd",
	}
)

func init() {
	for os, expression := range rootProcessExpressions {
		rootProcessRegexps[os] = regexp.MustCompile(expression)
	}
}

func GetRootProcess(os string, processes []*v1.Process) *v1.Process {
	var rootProcessRegexp *regexp.Regexp
	var found bool

	if rootProcessRegexp, found = rootProcessRegexps[os]; !found {
		return nil
	}

	for _, process := range processes {
		if rootProcessRegexp.Match([]byte(process.GetCommand())) {
			return process
		}
	}

	log.Printf("root process could not be found for OS %s", os)
	return nil
}

func PrintProcessLine(process string) {
	log.Printf("problematic process line: %s", process)
}

func PrintProcess(process *v1.Process) {
	log.Printf(
		"PID: %d\tPPID: %d\tChilds: %d\t Duration: %s\t=> %s (%d) => %s",
		process.GetPID(),
		process.GetPPID(),
		process.GetChilds(),
		chronos.DurationFromInt64(process.GetUptime()),
		process.GetUser(),
		process.GetUserID(),
		process.GetCommand(),
	)
}
