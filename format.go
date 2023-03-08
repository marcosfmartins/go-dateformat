package dateformat

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	shotNameSize = 3

	prefix     = "%"
	trimPrefix = prefix + "-"

	LongYear              = prefix + "Y"
	Year                  = prefix + "y"
	Month                 = prefix + "m"
	TrimMonth             = trimPrefix + "m"
	FullMonthName         = prefix + "B"
	ShortMonthName        = prefix + "b"
	FullWeekDayName       = prefix + "A"
	ShortWeekDayName      = prefix + "a"
	Day                   = prefix + "d"
	TrimDay               = trimPrefix + "d"
	DayYear               = prefix + "j"
	TrimDayYear           = trimPrefix + "j"
	WeekNumber            = prefix + "U"
	Hour24                = prefix + "H"
	Hour12                = prefix + "l"
	TrimHour12            = trimPrefix + "l"
	Minute                = prefix + "M"
	TrimMinute            = trimPrefix + "M"
	Second                = prefix + "S"
	TrimSecond            = trimPrefix + "S"
	Milliseconds          = prefix + "3f"
	MicroSeconds          = prefix + "6f"
	NanosecondsAlias      = prefix + "f"
	Nanoseconds           = prefix + "9f"
	AMPM                  = prefix + "P"
	AmPm                  = prefix + "p"
	TimezoneHH            = trimPrefix + "z"
	TimezoneHHMM          = prefix + "z"
	TimezoneWithColonHHMM = prefix + ":z"

	longYearSize     = 4
	monthSize        = 2
	daySize          = 2
	hour24Size       = 2
	hour12Size       = 2
	minuteSize       = 2
	secondSize       = 2
	milliSecondsSize = 3
	microSecondsSize = 6
	nanoSecondsSize  = 9

	weekDaysNumber = 7
	hour12Number   = 12
)

func Format(date time.Time, format string) string {
	reader := newReader(format)
	result := newStrBuilder()

	for {
		tok, err := reader.NextToken()
		if err != nil {
			break
		}

		switch tok {
		case LongYear:
			_, _ = result.WriteIntPrefix(date.Year(), longYearSize)
		case Year:
			str := strings.Split(strconv.Itoa(date.Year()), "")
			_, _ = result.WriteStringSlice(str[len(str)-2:])
		case Month:
			_, _ = result.WriteIntPrefix(int(date.Month()), monthSize)
		case TrimMonth:
			_ = result.WriteInt(int(date.Month()))
		case Day:
			_, _ = result.WriteIntPrefix(date.Day(), daySize)
		case TrimDay:
			_ = result.WriteInt(date.Day())
		case Hour24:
			_, _ = result.WriteIntPrefix(date.Hour(), hour24Size)
		case Hour12:
			_, _ = result.WriteIntPrefix(hourTo12(&date), hour12Size)
		case TrimHour12:
			_ = result.WriteInt(hourTo12(&date))
		case Minute:
			_, _ = result.WriteIntPrefix(date.Minute(), minuteSize)
		case TrimMinute:
			_ = result.WriteInt(date.Minute())
		case Second:
			_, _ = result.WriteIntPrefix(date.Second(), secondSize)
		case TrimSecond:
			_ = result.WriteInt(date.Second())
		case Milliseconds:
			_, _ = result.WriteString(fmt.Sprintf("%09d", date.Nanosecond())[:milliSecondsSize])
		case MicroSeconds:
			_, _ = result.WriteString(fmt.Sprintf("%09d", date.Nanosecond())[:microSecondsSize])
		case Nanoseconds, NanosecondsAlias:
			_, _ = result.WriteIntPrefix(date.Nanosecond(), nanoSecondsSize)
		case AMPM:
			_, _ = result.WriteString(getAmPm(&date))
		case AmPm:
			_, _ = result.WriteString(strings.ToLower(getAmPm(&date)))
		case FullMonthName:
			_, _ = result.WriteString(date.Month().String())
		case ShortMonthName:
			_, _ = result.WriteString(getShortName(date.Month().String()))
		case FullWeekDayName:
			_, _ = result.WriteString(date.Weekday().String())
		case ShortWeekDayName:
			_, _ = result.WriteString(getShortName(date.Weekday().String()))
		case DayYear:
			_, _ = result.WriteIntPrefix(date.YearDay(), 3)
		case TrimDayYear:
			_ = result.WriteInt(date.YearDay())
		case WeekNumber:
			_, _ = result.WriteString(fmt.Sprintf("%.0f", math.Ceil(float64(date.YearDay())/weekDaysNumber)))
		case TimezoneHHMM:
			_, offset := date.Zone()
			_, _ = result.WriteString(secondsHours(offset) + secondsMinuts(offset))
		case TimezoneWithColonHHMM:
			_, offset := date.Zone()
			_, _ = result.WriteString(secondsHours(offset) + ":" + secondsMinuts(offset))
		case TimezoneHH:
			_, offset := date.Zone()
			_, _ = result.WriteString(secondsHours(offset))
		default:
			if len(tok) > 1 {
				_, _ = result.WriteString("")
				continue
			}

			_, _ = result.WriteString(tok)
		}
	}

	return result.String()
}

func getShortName(s string) string {
	split := strings.Split(s, "")
	return strings.Join(split[:shotNameSize], "")
}

func secondsHours(input int) string {
	seconds := input % (60 * 60 * 24)
	hours := math.Floor(float64(seconds) / 60 / 60)

	signal := "+"
	if hours < 0 {
		signal = "-"
	}

	return fmt.Sprintf("%s%02d", signal, int(math.Abs(hours)))
}

func secondsMinuts(input int) string {
	seconds := input % (60 * 60)
	minutes := math.Floor(float64(seconds) / 60)

	return fmt.Sprintf("%02d", int(math.Abs(minutes)))
}

func hourTo12(datetime *time.Time) int {
	hour := datetime.Hour()
	if hour > hour12Number {
		hour = hour - hour12Number
	}
	return hour
}

func getAmPm(datetime *time.Time) string {
	if datetime.Hour()-12 >= 0 {
		return "PM"
	} else {
		return "AM"
	}
}
