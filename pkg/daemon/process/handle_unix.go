//go:build linux && darwin
//+build linux darwin

package process

import (
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	v1 "github.com/ylallemant/panopticon/pkg/api/v1"
	"github.com/ylallemant/panopticon/pkg/chronos"
	"github.com/ylallemant/panopticon/pkg/reporting"
)

const (
	time24Format   = "2006-January-_2 15:04 MST"
	timeDayFormat  = "2006-January-_2 3:04PM MST"
	timeWeekFormat = "2006-January-_2 Mon03PM MST"
	timeDateFormat = "Jan_22006 15:04"
	timeFormat     = "_2Jan06 15:04"
	timeTypeDay    = "DAY"
	timeTypeWeek   = "WEEK"
	timeTypeOld    = "OLD"
	timeNoon       = time.Duration(43200000000000)
)

var (
	numberRegexp   = regexp.MustCompile(`\d+`)
	spacesRegexp   = regexp.MustCompile(`\s+`)
	newlineRegexp  = regexp.MustCompile(`\n`)
	timeDayRegexp  = regexp.MustCompile(`(\d{1,2}):(\d{2})(PM|AM)`)
	time24Regexp   = regexp.MustCompile(`(\d{1,2}):(\d{2})`)
	timeDateRegexp = regexp.MustCompile(`([a-zA-Z]{3})(\d{1,2})`)
	timeWeekRegexp = regexp.MustCompile(`([a-zA-Z]{3})(\d{2})(PM|AM)`)
	timeRegexp     = regexp.MustCompile(`(\d{1,2})([a-zA-Z]{3})(\d{2})`)
	cmdIndex       = 0
)

func listProcesses() ([]*v1.Process, error) {
	out, err := exec.Command("ps", "-efw").Output()
	if err != nil {
		return nil, err
	}

	lines := newlineRegexp.Split(string(out), -1)
	list := []*v1.Process{}
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

	rootProcess := reporting.GetRootProcess(runtime.GOOS, list)
	if rootProcess == nil {
		log.Printf("WARN: root process could not be found for OS %s", runtime.GOOS)
		return list, nil
	}

	for _, process := range list {
		if process.GetPPID() <= rootProcess.GetPID() {
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

func processFromLine(line string) (*v1.Process, error) {
	line = sanitizeLine(line)

	if line == "" {
		return nil, nil
	}

	if strings.HasPrefix(line, "UID PID PPID") {
		return nil, nil
	}

	process := new(v1.Process)
	var err error

	// UID PID PPID C STIME TTY TIME CMD
	parts := strings.SplitN(line, " ", 8)

	if cmdIndex == 0 {
		cmdIndex = len(parts) - 1
	}

	PID, err := strconv.ParseInt(parts[1], 10, 32)
	if err != nil {
		reporting.PrintProcessLine(line)
		return nil, err
	}
	process.PID = int32(PID)

	PPID, err := strconv.ParseInt(parts[2], 10, 32)
	if err != nil {
		reporting.PrintProcessLine(line)
		return nil, err
	}
	process.PPID = int32(PPID)

	var processUser *user.User

	if numberRegexp.Match([]byte(parts[0])) {
		processUser, err = user.LookupId(parts[0])
		if err != nil {
			reporting.PrintProcessLine(line)
			return nil, err
		}
	} else {
		processUser, err = user.Lookup(parts[0])
		if err != nil {
			// TODO : ignore unknown users
			reporting.PrintProcessLine(line)
			log.Printf("WARN: user %s unknown", parts[0])
			return nil, nil
		}
	}

	UserID, err := strconv.ParseInt(processUser.Uid, 10, 32)
	if err != nil {
		reporting.PrintProcessLine(line)
		return nil, err
	}

	process.UserID = int32(UserID)
	process.User = processUser.Username

	startTime, err := parseTime(parts[4])
	if err != nil {
		reporting.PrintProcessLine(line)
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
	} else if time24Regexp.Match([]byte(raw)) {
		format = time24Format
		raw = fmt.Sprintf("%d-%s-%d %s %s", now.Year(), now.Month(), now.Day(), raw, zone)
	} else if timeWeekRegexp.Match([]byte(raw)) {
		format = timeWeekFormat
		weekday := raw[:3]
		day, err := chronos.SubWeekday(now, weekday)
		if err != nil {
			return time.Time{}, err
		}

		raw = fmt.Sprintf("%d-%s-%d %s %s", day.Year(), day.Month(), day.Day(), raw, zone)
	} else if timeDateRegexp.Match([]byte(raw)) {
		format = timeDateFormat
		raw = fmt.Sprintf("%s%d 12:00", raw, now.Year())
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
