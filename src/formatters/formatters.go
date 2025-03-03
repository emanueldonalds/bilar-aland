package formatters

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
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

func FormatPrice(value int32) string {
	if value <= 0 {
		return ""
	}
	var res = message.NewPrinter(language.Swedish).Sprintf("%d €", value)
	return res
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
