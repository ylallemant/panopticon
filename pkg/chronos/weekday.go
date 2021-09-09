package chronos

import (
	"fmt"
	"time"
)

var daysOfWeek = []string{}

func init() {
	for d := time.Sunday; d <= time.Saturday; d++ {
		name := d.String()
		daysOfWeek = append(daysOfWeek, name[:3])
	}
}

func SubWeekday(date time.Time, weekday string) (time.Time, error) {

	days, err := dayDiffrence(shortWeekday(date), weekday)
	if err != nil {
		return time.Time{}, err
	}

	duration, err := time.ParseDuration(fmt.Sprintf("%dh", days*24))
	if err != nil {
		return time.Time{}, err
	}

	past := date.Add(-duration)
	return past, nil
}

func dayDiffrence(today, weekday string) (int, error) {
	diff := 1

	index := indexOfWeekday(weekday) + 1
	if index < 0 {
		return -1, fmt.Errorf("could not find %s in weekdays", weekday)
	}

	overflow := 0

	for overflow < 2 {

		for index < len(daysOfWeek) {
			//log.Printf("     => %d\t%s <> %s\t%d", index, today, daysOfWeek[index], diff)
			if daysOfWeek[index] == today {
				return diff, nil
			}

			diff += 1
			index += 1
		}

		index = 0
		overflow += 1
	}

	if overflow > 1 {
		return -1, fmt.Errorf("could not find %s in weekdays", today)
	}

	return diff, nil
}

func indexOfWeekday(weekday string) int {
	for index, day := range daysOfWeek {
		if day == weekday {
			return index
		}
	}

	return -1
}

func shortWeekday(weekday time.Time) string {
	return weekday.Weekday().String()[:3]
}
