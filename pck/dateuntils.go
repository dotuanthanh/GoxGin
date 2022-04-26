package pck

import (
	"time"
)

const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"

	RFC3339X = "2006-01-02 15:04:05"
)

func ConvertUnixTimeToUTC(t int64) time.Time {
	unixTimeUTC := time.Unix(t, 0) //gives unix time stamp in utc
	unixTimeUTC.Format(RFC3339X)
	return unixTimeUTC

}

func ConvertUnixTimeToRFC3339(t int64) string {
	unixTimeUTC := time.Unix(t, 0) //gives unix time stamp in utc

	unitTimeInRFC3339 := unixTimeUTC.Format(time.RFC3339) // converts utc time to RFC3339 format
	return unitTimeInRFC3339

}

func GetTimeStp() string {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	t := time.Now().In(loc)
	return t.Format("20060102150405")
}
func GetTodaysDate() string {
	loc, _ := time.LoadLocation("Tokyo")
	current_time := time.Now().In(loc)
	return current_time.Format("2006-01-02")
}

func GetTodaysDateTime() string {
	loc, _ := time.LoadLocation("Tokyo")
	current_time := time.Now().In(loc)
	return current_time.Format("2006-01-02 15:04:05")
}

func GetTodaysDateTimeFormatted() string {
	loc, _ := time.LoadLocation("Tokyo")
	current_time := time.Now().In(loc)
	return current_time.Format("Jan 2, 2006 at 3:04 PM")
}

func GetTimeStampFromDate(dtformat string) string {
	form := "Jan 2, 2006 at 3:04 PM"
	t2, _ := time.Parse(form, dtformat)
	return t2.Format("20060102150405")
}
func FormatDate(num string, layout string) time.Time {
	parse, err := time.Parse("2006-01-02", num)
	if err != nil {
		return time.Time{}
	}
	return parse
}
func GetTimeStamp() int64 {
	return time.Now().Unix()
}
