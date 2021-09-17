//go:build linux && darwin
//+build linux darwin

package process

import (
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"regexp"
	"strconv"
	"strings"
	"time"

	api "github.com/ylallemant/panopticon/pkg/api/v1"
	"github.com/ylallemant/panopticon/pkg/chronos"
)

const (
	timeDayFormat  = "2006-January-_2 3:04PM MST"
	timeWeekFormat = "2006-January-_2 Mon03PM MST"
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
	log.Printf("%d lines found", len(lines))

	for _, line := range lines {
		process, err := processFromLine(line)
		if err != nil {
			return nil, err
		}

		if process != nil {
			//log.Printf("line %d\t=> PID %d\t=> User %d (%s)", idx, process.PID, process.UserID, process.User)
			list = append(list, process)
		}
	}

	for _, process := range list {
		if process.GetPPID() < 2 {
			childs := 0

			for _, current := range list {
				if process.GetPID() == current.GetPPID() {
					childs += 1
				}
			}

			process.Childs = int32(childs)
		}
	}

	log.Printf("%d lines parsed", len(list))
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

	PID, err := strconv.ParseInt(parts[1], 10, 32)
	if err != nil {
		return nil, err
	}
	process.PID = int32(PID)

	PPID, err := strconv.ParseInt(parts[2], 10, 32)
	if err != nil {
		return nil, err
	}
	process.PPID = int32(PPID)

	UserID, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return nil, err
	}
	process.UserID = int32(UserID)

	processUser, err := user.LookupId(parts[0])
	if err != nil {
		return nil, err
	}

	process.User = processUser.Username

	startTime, err := parseTime(parts[4])
	if err != nil {
		return nil, err
	}

	//log.Printf("\t\tstart: %s", startTime)
	now := time.Now().UTC()
	//log.Printf("\t\t  now: %s", now)
	uptime := now.Sub(startTime)
	//log.Printf("\t\t diff: %s", uptime)

	process.Uptime = chronos.DurationToNanoseconds(uptime)
	//log.Printf("%s\tuptime = %s \t\t=>%d\n", parts[4], uptime, process.Uptime)

	process.Command = parts[cmdIndex]
	//log.Printf("\t%+#v\n", process)

	return process, nil
}

func parseTime(raw string) (time.Time, error) {
	format := ""
	now := time.Now()
	//trueRaw := raw
	zone, _ := now.Zone()

	if timeDayRegexp.Match([]byte(raw)) {
		format = timeDayFormat
		raw = fmt.Sprintf("%d-%s-%d %s %s", now.Year(), now.Month(), now.Day(), raw, zone)
	} else if timeWeekRegexp.Match([]byte(raw)) {
		format = timeWeekFormat
		weekday := raw[:3]
		day, err := chronos.SubWeekday(now, weekday)
		if err != nil {
			return time.Time{}, err
		}

		raw = fmt.Sprintf("%d-%s-%d %s %s", day.Year(), day.Month(), day.Day(), raw, zone)
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
	return parsed.UTC(), nil
}
