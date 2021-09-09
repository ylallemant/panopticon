package process

// +build linux
// +build darwin

import (
	"fmt"
	"os/exec"
	"os/user"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ylallemant/panopticon/pkg/api"
	"github.com/ylallemant/panopticon/pkg/chronos"
)

const (
	timeDayFormat  = "2006-January-_2 3:04PM"
	timeWeekFormat = "2006-January-_2 Mon03PM"
	timeFormat     = "_2Jan06 15:04"
	timeTypeDay    = "DAY"
	timeTypeWeek   = "WEEK"
	timeTypeOld    = "OLD"
	timeNoon       = time.Duration(43200000000000)
)

var (
	spacesRegexp   = regexp.MustCompile(`\s+`)
	newlineRegexp  = regexp.MustCompile(`\n`)
	timeDayRegexp  = regexp.MustCompile(`(\d{1,2}):(\d{2})(PM|AM)`)
	timeWeekRegexp = regexp.MustCompile(`([a-zA-Z]{3})(\d{2})(PM|AM)`)
	timeRegexp     = regexp.MustCompile(`(\d{1,2})([a-zA-Z]{3})(\d{2})`)
	cmdIndex       = 0
)

func listProcesses() ([]*api.Process, error) {
	out, err := exec.Command("ps", "-efw").Output()
	if err != nil {
		return nil, err
	}

	lines := newlineRegexp.Split(string(out), -1)
	list := []*api.Process{}

	for _, line := range lines {
		process, err := processFromLine(line)
		if err != nil {
			return nil, err
		}

		if process != nil {
			list = append(list, process)
		}
	}

	return list, nil
}

func sanitizeLine(line string) string {
	line = spacesRegexp.ReplaceAllString(line, " ")
	line = strings.TrimSpace(line)
	return line
}

func processFromLine(line string) (*api.Process, error) {
	line = sanitizeLine(line)

	if line == "" {
		return nil, nil
	}

	if strings.HasPrefix(line, "UID PID PPID") {
		return nil, nil
	}
	//log.Printf("\t%s\n", line)

	process := new(api.Process)
	var err error

	// UID PID PPID C STIME TTY TIME CMD
	parts := strings.SplitN(line, " ", 8)

	if cmdIndex == 0 {
		cmdIndex = len(parts) - 1
	}

	process.PID, err = strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	process.PPID, err = strconv.Atoi(parts[2])
	if err != nil {
		return nil, err
	}

	process.UserID, err = strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	processUser, err := user.LookupId(parts[0])
	if err != nil {
		return nil, err
	}

	process.User = processUser.Username

	startTime, err := parseTime(parts[4])
	if err != nil {
		return nil, err
	}

	process.StartTime = startTime

	process.Command = parts[cmdIndex]
	//log.Printf("\t%+#v\n", process.StartTime.String())

	return process, nil
}

func parseTime(raw string) (time.Time, error) {
	format := ""
	now := time.Now()

	if timeDayRegexp.Match([]byte(raw)) {
		format = timeDayFormat
		raw = fmt.Sprintf("%d-%s-%d %s", now.Year(), now.Month(), now.Day(), raw)
	} else if timeWeekRegexp.Match([]byte(raw)) {
		format = timeWeekFormat
		weekday := raw[:3]
		day, err := chronos.SubWeekday(now, weekday)
		if err != nil {
			return time.Time{}, err
		}

		raw = fmt.Sprintf("%d-%s-%d %s", day.Year(), day.Month(), day.Day(), raw)
	} else if timeRegexp.Match([]byte(raw)) {
		format = timeFormat
		raw = fmt.Sprintf("%s 12:00", raw)
	}

	parsed, err := time.Parse(format, raw)
	if err != nil {
		//log.Printf("ERROR\t%s\t=> %s\t=> %s\n", raw, format, err.Error())
		return time.Time{}, err
	}

	//log.Printf("   %s => %s\t\t=> %s (%s)", trueRaw, raw, parsed, format)
	return parsed, nil
}
