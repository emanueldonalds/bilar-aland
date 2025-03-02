package formatters

import (
	"fmt"
	"time"
)

var swedishMonths map[time.Month]string = map[time.Month]string{
	time.January:   "Jan",
	time.February:  "Feb",
	time.March:     "Mar",
	time.April:     "Apr",
	time.May:       "Maj",
	time.June:      "Jun",
	time.July:      "Jul",
	time.August:    "Aug",
	time.September: "Sep",
	time.October:   "Okt",
	time.November:  "Nov",
	time.December:  "Dec",
}

func FormatPrice(value string) string {
	if value == "" {
		return ""
	}
	return value + " €"
}

func FormatDateTime(value string) string {
	if value == "" {
		return ""
	}
	t := parseTime(value)
	formatted := fmt.Sprintf("%d %s %02d:%02d", t.Day(), swedishMonths[t.Month()], t.Hour(), t.Minute())
	return formatted
}

func FormatDateTimeRfc822(value string) string {
	if value == "" {
		return ""
	}
	t := parseTime(value)
	formatted := t.Format("Mon, 02 Jan 2006 15:04:05 -0700")
	return formatted
}

func parseTime(value string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05.999999", value)
	if err != nil {
		panic(err.Error())
	}
	return t.In(time.Local)
}
