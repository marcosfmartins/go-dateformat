package dateformat

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDateFormat(t *testing.T) {
	datetime := time.Now()

	DateOnly := "2006-01-02"
	TimeOnly := "15:04:05"
	DateTime := "2006-01-02 15:04:05"
	UnixDate := "Mon Jan _2 15:04:05 -07 2006"
	RFC850 := "Monday, 02-Jan-06 15:04:05 -07"
	RFC1123 := "Mon, 02 Jan 2006 15:04:05 -07"
	RFC3339 := "2006-01-02T15:04:05Z-07:00"
	RFC3339Nano := "2006-01-02T15:04:05.999999999Z-07:00"

	models := []struct {
		Name     string
		Expected string
		ToTest   string
	}{
		{
			Name:     "Layout",
			Expected: time.Layout,
			ToTest:   "%m/%d %l:%M:%S%P '%y %z",
		},
		{
			Name:     "ANSIC",
			Expected: time.ANSIC,
			ToTest:   "%a %b  %-d %H:%M:%S %Y",
		},
		{
			Name:     "UnixDate",
			Expected: UnixDate,
			ToTest:   "%a %b  %-d %H:%M:%S %-z %Y",
		},
		{
			Name:     "RubyDate",
			Expected: time.RubyDate,
			ToTest:   "%a %b %d %H:%M:%S %z %Y",
		},
		{
			Name:     "RFC822Z",
			Expected: time.RFC822Z,
			ToTest:   "%d %b %y %H:%M %z",
		},
		{
			Name:     "RFC850",
			Expected: RFC850,
			ToTest:   "%A, %d-%b-%y %H:%M:%S %-z",
		},
		{
			Name:     "RFC1123",
			Expected: RFC1123,
			ToTest:   "%a, %d %b %Y %H:%M:%S %-z",
		},
		{
			Name:     "RFC1123Z",
			Expected: time.RFC1123Z,
			ToTest:   "%a, %d %b %Y %H:%M:%S %z",
		},
		{
			Name:     "RFC3339",
			Expected: RFC3339,
			ToTest:   "%Y-%m-%dT%H:%M:%SZ%:z",
		},
		{
			Name:     "RFC3339Nano",
			Expected: RFC3339Nano,
			ToTest:   "%Y-%m-%dT%H:%M:%S.%fZ%:z",
		},
		{
			Name:     "Kitchen",
			Expected: time.Kitchen,
			ToTest:   "%-l:%M%P",
		},
		{
			Name:     "Stamp",
			Expected: time.Stamp,
			ToTest:   "%b  %-d %H:%M:%S",
		},
		{
			Name:     "StampMilli",
			Expected: time.StampMilli,
			ToTest:   "%b  %-d %H:%M:%S.%3f",
		},
		{
			Name:     "StampMicro",
			Expected: time.StampMicro,
			ToTest:   "%b  %-d %H:%M:%S.%6f",
		},
		{
			Name:     "StampNano",
			Expected: time.StampNano,
			ToTest:   "%b  %-d %H:%M:%S.%f",
		},
		{
			Name:     "DateTime",
			Expected: DateTime,
			ToTest:   "%Y-%m-%d %H:%M:%S",
		},
		{
			Name:     "DateOnly",
			Expected: DateOnly,
			ToTest:   "%Y-%m-%d",
		},
		{

			Name:     "TimeOnly",
			Expected: TimeOnly,
			ToTest:   "%H:%M:%S",
		},
		{
			Name:     "invalid",
			Expected: "  ",
			ToTest:   "%3x %-k %:r",
		},
	}

	for _, model := range models {
		t.Run(model.Name, func(t *testing.T) {
			expected := datetime.Format(model.Expected)
			result := Format(datetime, model.ToTest)
			assert.Equal(t, expected, result)
		})
	}
}
